package consumers

import (
	"context"
	"fmt"
	"log"

	"github.com/xrpscan/platform/connections"
)

func RunDefaultConsumer() {
	go func() {
		ctx := context.Background()
		r := connections.KafkaReader
		for {
			m, err := r.FetchMessage(ctx)
			if err != nil {
				break
			}
			fmt.Printf("Message at topic(%v), partition(%v), offset(%v): %s\n", m.Topic, m.Partition, m.Offset, string(m.Key))
			// processor.IndexTransaction(m.Key, m.Value)

			// Explicitly commit kafka message
			if err := r.CommitMessages(ctx, m); err != nil {
				log.Println("Failed to commit messages to kafka: ", err)
			}
		}
	}()
}
