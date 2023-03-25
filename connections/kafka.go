package connections

import (
	"sync"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/config"
)

var KafkaWriter *kafka.Writer
var wOnce sync.Once

func NewWriter() {
	wOnce.Do(func() {
		KafkaWriter = &kafka.Writer{
			Addr:     kafka.TCP(config.EnvKafkaBootstrapServer()),
			Balancer: &kafka.LeastBytes{},
			Topic:    "test.messages",
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
			Topic:   "test.messages",
		})
	})
}
