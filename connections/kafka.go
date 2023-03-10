package connections

import (
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var KafkaProducer *kafka.Producer
var once_kp sync.Once

func NewProducer() {
	once_kp.Do(func() {
		kp, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost",
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
			"bootstrap.servers": "localhost",
			"group.id":          "test.dragons",
		})
		if err != nil {
			panic(err)
		}
		KafkaConsumer = kc
	})
}
