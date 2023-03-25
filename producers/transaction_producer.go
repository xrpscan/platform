package producers

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func Produce(w *kafka.Writer, message string) {
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("42"),
			Value: []byte(message),
		},
	)
	if err != nil {
		log.Printf("Failed to write message: %s", err)
	}
}
