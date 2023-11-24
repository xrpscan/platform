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
func IndexValidation(m kafka.Message) {
	logger.Log.Debug().Str("topic", m.Topic).Int("partition", m.Partition).Int64("offset", m.Offset).Str("key", string(m.Key)).Msg("Indexing")

	key, message := m.Key, m.Value
	var validation map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(string(message))).Decode(&validation); err != nil {
		logger.Log.Error().Err(err).Msg("Error decoding validation json")
		return
	}

	req := esapi.IndexRequest{
		Index:      models.StreamValidation.String(),
		DocumentID: string(key),
		Body:       strings.NewReader(string(message)),
	}
	Index(req)
}

func BulkIndexValidation(ch <-chan kafka.Message) {
	bulk, _ := NewBulkIndexClient(models.StreamValidation.String())

	// Bulk index channel reader loop
	for {
		message := <-ch

		ledgerIndex, err := GetLedgerIndex(message.Value)
		if err != nil {
			logger.Log.Debug().Err(err).Msg("ledger_index not found in validation")
			continue
		}
		indexName := GetIndexName(models.StreamValidation.String(), ledgerIndex)

		err = bulk.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action:     "index",
				Index:      indexName,
				DocumentID: string(message.Key),
				Body:       bytes.NewReader(message.Value),
			},
		)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error adding documents to bulk indexer")
		}
	}
}
