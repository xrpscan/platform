package processor

import (
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/xrpl/models"
)

func PrintMessage(m kafka.Message) {
	fmt.Printf("[PrintMessage] Message at topic(%v), partition(%v), offset(%v): %s\n", m.Topic, m.Partition, m.Offset, string(m.Key))
}

func Test(m kafka.Message) {
	fmt.Printf("[Test] Message at topic(%v), partition(%v), offset(%v): %s\n", m.Topic, m.Partition, m.Offset, string(m.Key))

	req := models.AccountInfoRequest{
		BaseRequest: models.BaseRequest{Id: connections.XrplClient.NextID(), Command: "account_info"},
		Account:     "rw2ciyaNshpHe7bCHo4bRWq6pqqynnWKQg",
	}
	j, _ := json.Marshal(req)
	fmt.Println(string(j))
	connections.XrplClient.Request(j)
}
