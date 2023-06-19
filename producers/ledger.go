package producers

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/xrpl-go"
)

func ProduceLedger(w *kafka.Writer, message []byte) {
	var res xrpl.BaseResponse
	if err := json.Unmarshal(message, &res); err != nil {
		return
	}

	messageKey := strconv.Itoa(int(res["ledger_index"].(float64)))

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Topic: config.TopicLedgers(),
			Key:   []byte(messageKey),
			Value: message,
		},
	)
	if err != nil {
		log.Printf("Failed to write message: %s", err)
	}
}
