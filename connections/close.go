package connections

import "log"

func CloseAll() {
	log.Println("Closing all connections")
	KafkaReader.Close()
	KafkaWriter.Close()
}
