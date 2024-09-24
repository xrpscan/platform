package indexer

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/platform/models"
)

// Deprecated: This function does not support new IndexTemplate based indexing
func IndexTransaction(m kafka.Message) {
	logger.Log.Debug().Str("topic", m.Topic).Int("partition", m.Partition).Int64("offset", m.Offset).Str("key", string(m.Key)).Msg("Indexing")

	key, message := m.Key, m.Value
	var tx map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(string(message))).Decode(&tx); err != nil {
		logger.Log.Error().Err(err).Msg("Error decoding transaction json")
		return
	}

	req := esapi.IndexRequest{
		Index:      models.StreamTransaction.String(),
		DocumentID: string(key),
		Body:       strings.NewReader(string(message)),
	}
	Index(req)
}

func BulkIndexTransaction(ch <-chan kafka.Message) {
	bulk, _ := NewBulkIndexClient(models.StreamTransaction.String())

	// Kafka message reader loop
	for {
		message := <-ch

		// Unmarshal message.Value to MapStringInterface for further processing
		var tx map[string]interface{}
		err := json.Unmarshal(message.Value, &tx)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Transaction json.Unmarshal error")
			continue
		}

		// Modify transaction by fixing some transaction fields
		modifiedTx, err := ModifyTransaction(tx)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error fixing transaction object")
			continue
		}

		// Marshal modified transaction back to JSON
		txJSON, err := json.Marshal(modifiedTx)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Transaction json.Marshal error")
			continue
		}

		li, ok := tx["ledger_index"].(float64)
		if !ok {
			logger.Log.Error().Err(err).Msg("ledger_index not found in transaction")
			continue
		}
		indexName := GetIndexName(models.StreamTransaction.String(), int(li))

		// Add tx to bulk indexing queue
		err = bulk.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action:     "index",
				Index:      indexName,
				DocumentID: string(message.Key),
				Body:       bytes.NewReader(txJSON),
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						logger.Log.Error().Err(err).Msg("Bulk index error")
					} else {
						logger.Log.Error().Err(err).Str("hash", item.DocumentID).Str("type", res.Error.Type).Str("reason", res.Error.Reason).Msg("Bulk index error")
					}
				},
			},
		)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error adding documents to bulk indexer")
		}
	}
}
