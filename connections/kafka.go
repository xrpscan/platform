package connections

import (
	"sync"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/config"
)

var KafkaWriter *kafka.Writer
var wOnce sync.Once

// Common Kafka writer for the application. Every message must
// specify the Topic where it must be written to.
func NewWriter() {
	wOnce.Do(func() {
		KafkaWriter = &kafka.Writer{
			Addr:     kafka.TCP(config.EnvKafkaBootstrapServer()),
			Balancer: &kafka.LeastBytes{},
			Async:    true,
		}
	})
}

var KafkaReader *kafka.Reader
var rOnce sync.Once

func NewReader() {
	rOnce.Do(func() {
		KafkaReader = kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{config.EnvKafkaBootstrapServer()},
			GroupID: config.EnvKafkaGroupId(),
			Topic:   config.TopicTransactions(),
		})
	})
}
