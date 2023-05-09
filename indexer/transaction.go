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

		err := bulk.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: string(message.Key),
				Body:       bytes.NewReader(message.Value),
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						logger.Log.Trace().Err(err).Msg("Bulk index error")
					} else {
						logger.Log.Trace().Err(err).Str("hash", item.DocumentID).Str("type", res.Error.Type).Str("reason", res.Error.Reason).Msg("Bulk index error")
					}
				},
			},
		)
		if err != nil {
			logger.Log.Error().Err(err).Msg("Error adding documents to bulk indexer")
		}
	}
}
