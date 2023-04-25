package producers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/xrpl"
)

func ProduceTransaction(w *kafka.Writer, message []byte) {
	var res xrpl.BaseResponse
	if err := json.Unmarshal(message, &res); err != nil {
		fmt.Println("json.Unmarshal error: ", err)
	}

	tx, ok := res["transaction"].(map[string]interface{})
	if !ok {
		return
	}
	messageKey, ok := tx["hash"].(string)
	if !ok {
		return
	}

	err := w.WriteMessages(context.Background(),
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
