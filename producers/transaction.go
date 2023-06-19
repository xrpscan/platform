package producers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/xrpl-go"
)

func ProduceTransaction(w *kafka.Writer, message []byte) {
	var res xrpl.BaseResponse
	err := json.Unmarshal(message, &res)
	if err != nil {
		return
	}

	tx, ok := res["transaction"].(map[string]interface{})
	if !ok {
		return
	}
	messageKey, ok := tx["hash"].(string)
	if !ok {
		return
	}

	err = w.WriteMessages(context.Background(),
		kafka.Message{
			Topic: config.TopicTransactions(),
			Key:   []byte(messageKey),
			Value: message,
		},
	)
	if err != nil {
		log.Printf("Failed to write message: %s", err)
	}
}
