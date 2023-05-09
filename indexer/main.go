package indexer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/platform/xrpl"
)

func PrintMessage(m kafka.Message) {
	msg := fmt.Sprintf("Message on topic(%v), partition(%v), offset(%v): %s", m.Topic, m.Partition, m.Offset, string(m.Key))
	logger.Log.Info().Msg(msg)
}

func Test(m kafka.Message) {
	fmt.Printf("[Test] Message at topic(%v), partition(%v), offset(%v): %s\n", m.Topic, m.Partition, m.Offset, string(m.Key))
	req := xrpl.BaseRequest{
		"command": "account_info",
		"account": "rw2ciyaNshpHe7bCHo4bRWq6pqqynnWKQg",
	}
	res, _ := connections.XrplClient.Request(req)
	fmt.Println("Response:", res)
}

func Index(req esapi.IndexRequest) {
	ctx := context.Background()
	res, err := req.Do(ctx, connections.GetEsClient())
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error indexing document")
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			logger.Log.Error().Err(err).Msg("Error decoding index error message")
		}
	}
}

// Create a new Elasticsearch bulk index client
func NewBulkIndexClient(index string) (esutil.BulkIndexer, error) {
	flushInterval := 5 * time.Second
	bulk, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         index,
		Client:        connections.EsClient,
		NumWorkers:    1,
		FlushBytes:    int(1024 * 1024),
		FlushInterval: flushInterval,
	})
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error creating bulk indexer")
		return nil, err
	}

	// Log bulk index stats
	go func() {
		for range time.Tick(flushInterval) {
			bs := bulk.Stats()
			logger.Log.Debug().Str("index", index).Uint64("indexed", bs.NumIndexed).Uint64("flushed", bs.NumFlushed).Uint64("added", bs.NumAdded).Uint64("reqs", bs.NumRequests).Uint64("failed", bs.NumFailed).Msg("Index bulk")
		}
	}()

	return bulk, nil
}
