package connections

import (
	"log"
)

func CloseWriter() {
	if err := KafkaWriter.Close(); err != nil {
		log.Println("Error closing Kafka writer connection")
	}
}

func CloseReaders() {
	if err := KafkaReaderLedger.Close(); err != nil {
		log.Println("Error closing Kafka Ledger reader connection")
	}

	if err := KafkaReaderTransaction.Close(); err != nil {
		log.Println("Error closing Kafka Transaction reader connection")
	}

	if err := KafkaReaderValidation.Close(); err != nil {
		log.Println("Error closing Kafka Validation reader connection")
	}

	if err := KafkaReaderPeerStatus.Close(); err != nil {
		log.Println("Error closing Kafka PeerStatus reader connection")
	}

	if err := KafkaReaderConsensus.Close(); err != nil {
		log.Println("Error closing Kafka Consensus reader connection")
	}

	if err := KafkaReaderPathFind.Close(); err != nil {
		log.Println("Error closing Kafka PathFind reader connection")
	}

	if err := KafkaReaderManifest.Close(); err != nil {
		log.Println("Error closing Kafka Manifest reader connection")
	}

	if err := KafkaReaderServer.Close(); err != nil {
		log.Println("Error closing Kafka Server reader connection")
	}

	if err := KafkaReaderDefault.Close(); err != nil {
		log.Println("Error closing Kafka Default reader connection")
	}
}

func CloseEsClient() {
}

func CloseXrplClient() {
	if err := XrplClient.Close(); err != nil {
		log.Println("Error closing xrpl connection")
	}
}

func CloseAll() {
	log.Println("Closing all connections")
	CloseWriter()
	CloseReaders()
	CloseEsClient()
	CloseXrplClient()
}
