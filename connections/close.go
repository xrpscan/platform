package connections

import (
	"log"
)

func CloseAll() {
	log.Println("Closing all connections")
	if err := KafkaReader.Close(); err != nil {
		log.Println("Error closing Kafka reader connection")
	}

	if err := KafkaWriter.Close(); err != nil {
		log.Println("Error closing Kafka writer connection")
	}

	if err := XrplClient.Close(); err != nil {
		log.Println("Error closing xrpl connection")
	}
}
