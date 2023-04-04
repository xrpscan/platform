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

	if XrplClient != nil && XrplClient.IsConnected() {
		err := XrplClient.Close()
		if err != nil {
			log.Panicln("Error closing xrpl connection")
		}
	}

	if XrplFHClient != nil && XrplFHClient.IsConnected() {
		err := XrplFHClient.Close()
		if err != nil {
			log.Panicln("Error closing xrpl fullhistory connection")
		}
	}
}
