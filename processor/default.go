package processor

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

func PrintMessage(m kafka.Message) {
	fmt.Printf("Message at topic(%v), partition(%v), offset(%v): %s\n", m.Topic, m.Partition, m.Offset, string(m.Key))
}
