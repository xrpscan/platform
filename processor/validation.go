package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/connections"
)

func IndexValidation(m kafka.Message) {
	key, message := m.Key, m.Value
	var validation map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(string(message))).Decode(&validation); err != nil {
		fmt.Println("Error decoding validation")
	}

	req := esapi.IndexRequest{
		Index:      "validation",
		DocumentID: string(key),
		Body:       strings.NewReader(string(message)),
	}

	ctx := context.Background()
	res, err := req.Do(ctx, connections.GetEsClient())
	if err != nil {
		fmt.Println("Error indexing document: " + string(key))
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			fmt.Println("Error decoding index error message")
		}
		fmt.Printf("[%s] %s: %s\n", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
	}
}
