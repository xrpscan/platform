package indexer

import (
	"encoding/json"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/platform/models"
)

func IndexLedger(m kafka.Message) {
	logger.Log.Debug().Str("topic", m.Topic).Int("partition", m.Partition).Int64("offset", m.Offset).Str("key", string(m.Key)).Msg("Indexing")

	key, message := m.Key, m.Value
	var ledger map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(string(message))).Decode(&ledger); err != nil {
		logger.Log.Error().Err(err).Msg("Error decoding ledger json")
		return
	}

	req := esapi.IndexRequest{
		Index:      models.StreamLedger.String(),
		DocumentID: string(key),
		Body:       strings.NewReader(string(message)),
	}
	Index(req)
}
