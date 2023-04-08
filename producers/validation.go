package producers

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/config"
)

func ProduceValidation(w *kafka.Writer, message []byte) {
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Topic: config.TopicValidations(),
			Key:   []byte(uuid.New().String()),
			Value: message,
		},
	)
	if err != nil {
		log.Printf("Failed to write message: %s", err)
	}
}
