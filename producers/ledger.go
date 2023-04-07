package producers

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/xrpl"
)

func SubscribeStreams() {
	connections.XrplClient.Subscribe("ledger")
	connections.XrplClient.Subscribe("transactions")

	go func() {
		for {
			select {
			case ledger := <-connections.XrplClient.LedgerStream:
				Produce(connections.KafkaWriter, ledger)

			case validation := <-connections.XrplClient.ValidationStream:
				Produce(connections.KafkaWriter, validation)

			case transaction := <-connections.XrplClient.TransactionStream:
				Produce(connections.KafkaWriter, transaction)

			case peerStatus := <-connections.XrplClient.PeerStatusStream:
				Produce(connections.KafkaWriter, peerStatus)

			case consensus := <-connections.XrplClient.ValidationStream:
				Produce(connections.KafkaWriter, consensus)

			case pathFind := <-connections.XrplClient.PeerStatusStream:
				Produce(connections.KafkaWriter, pathFind)

			case defaultObject := <-connections.XrplClient.DefaultStream:
				Produce(connections.KafkaWriter, defaultObject)
			}
		}
	}()
}

func Produce(w *kafka.Writer, message xrpl.StreamMessage) {
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   message.Key,
			Value: message.Value,
		},
	)
	if err != nil {
		log.Printf("Failed to write message: %s", err)
	}
}
