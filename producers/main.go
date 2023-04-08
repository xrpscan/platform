package producers

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/connections"
)

func SubscribeStreams() {
	// connections.XrplClient.Subscribe([]byte("ledger"))
	// connections.XrplClient.Subscribe([]byte("transactions"))
	// connections.XrplClient.Subscribe([]byte("validations"))

	go func() {
		for {
			select {
			case ledger := <-connections.XrplClient.StreamLedger:
				ProduceLedger(connections.KafkaWriter, ledger)

			case validation := <-connections.XrplClient.StreamValidation:
				ProduceValidation(connections.KafkaWriter, validation)

			case transaction := <-connections.XrplClient.StreamTransaction:
				ProduceTransaction(connections.KafkaWriter, transaction)

			case peerStatus := <-connections.XrplClient.StreamPeerStatus:
				Produce(connections.KafkaWriter, peerStatus)

			case consensus := <-connections.XrplClient.StreamConsensus:
				Produce(connections.KafkaWriter, consensus)

			case pathFind := <-connections.XrplClient.StreamPathFind:
				Produce(connections.KafkaWriter, pathFind)

			case manifest := <-connections.XrplClient.StreamManifest:
				Produce(connections.KafkaWriter, manifest)

			case server := <-connections.XrplClient.StreamServer:
				Produce(connections.KafkaWriter, server)

			case defaultObject := <-connections.XrplClient.StreamDefault:
				fmt.Println(string(defaultObject))
				Produce(connections.KafkaWriter, defaultObject)
			}
		}
	}()
}

func Produce(w *kafka.Writer, message []byte) {
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Topic: config.TopicDefault(),
			Key:   []byte(uuid.New().String()),
			Value: message,
		},
	)
	if err != nil {
		log.Printf("Failed to write message: %s", err)
	}
}
