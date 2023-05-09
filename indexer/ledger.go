package indexer

import (
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/platform/models"
)

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
