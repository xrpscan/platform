package indexer

import (
	"bytes"
	"context"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/platform/models"
)

// Deprecated: This function does not support new IndexTemplate based indexing
func IndexLedger(m kafka.Message) {
	key, message := m.Key, m.Value
	logger.Log.Debug().Str("index", models.StreamLedger.String()).Int("partition", m.Partition).Int64("offset", m.Offset).Str("key", string(key)).Msg("Index serial")

	req := esapi.IndexRequest{
		Index:      models.StreamLedger.String(),
		DocumentID: string(key),
		Body:       strings.NewReader(string(message)),
	}
	Index(req)
}

func BulkIndexLedger(ch <-chan kafka.Message) {
	bulk, _ := NewBulkIndexClient(models.StreamLedger.String())

	// Kafka message reader loop
	for {
		message := <-ch

		ledgerIndex, err := GetLedgerIndex(message.Value)
		if err != nil {
			logger.Log.Debug().Err(err).Msg("ledger_index not found in ledger")
			continue
		}
		indexName := GetIndexName(models.StreamLedger.String(), ledgerIndex)

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
