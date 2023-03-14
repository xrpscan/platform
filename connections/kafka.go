package connections

import (
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/xrpscan/platform/config"
)

var KafkaProducer *kafka.Producer
var once_kp sync.Once

func NewProducer() {
	once_kp.Do(func() {
		kp, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": config.EnvKafkaBootstrapServer(),
		})
		if err != nil {
			panic(err)
		}
		KafkaProducer = kp
	})
}

var KafkaConsumer *kafka.Consumer
var once_kc sync.Once

func NewConsumer() {
	once_kc.Do(func() {
		kc, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": config.EnvKafkaBootstrapServer(),
			"group.id":          config.EnvKafkaGroupId(),
		})
		if err != nil {
			panic(err)
		}
		KafkaConsumer = kc
	})
}
