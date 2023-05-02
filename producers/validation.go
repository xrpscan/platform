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

func ProduceValidation(w *kafka.Writer, message []byte) {
	var res xrpl.BaseResponse
	if err := json.Unmarshal(message, &res); err != nil {
		fmt.Println("json.Unmarshal error: ", err)
	}

	messageKey := fmt.Sprintf("%s.%s", res["validation_public_key"], res["ledger_index"])

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Topic: config.TopicValidations(),
			Key:   []byte(messageKey),
			Value: message,
		},
	)
	if err != nil {
		log.Printf("Failed to write message: %s", err)
	}
}
