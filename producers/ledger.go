package producers

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/connections"
)

func SubscribeStreams() {
	connections.XrplClient.Subscribe("ledger")
	connections.XrplClient.Subscribe("transactions")

	go func() {
		for {
			select {
			case ledger := <-connections.XrplClient.LedgerStream:
				Produce(connections.KafkaWriter, ledger)
			case tx := <-connections.XrplClient.TransactionStream:
				Produce(connections.KafkaWriter, tx)
			}
		}
	}()
}

func Produce(w *kafka.Writer, message []byte) {
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("42"),
			Value: message,
		},
	)
	if err != nil {
		log.Printf("Failed to write message: %s", err)
	}
}
