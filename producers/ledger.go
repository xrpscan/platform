package producers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/xrpl"
)

func ProduceLedger(w *kafka.Writer, message []byte) {
	var res xrpl.BaseResponse
	if err := json.Unmarshal(message, &res); err != nil {
		fmt.Println("json.Unmarshal error: ", err)
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
