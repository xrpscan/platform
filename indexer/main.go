package indexer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/esapi"
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
