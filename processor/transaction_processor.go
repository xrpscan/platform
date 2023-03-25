package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/xrpscan/platform/connections"
)

func IndexTransaction(tx string) {
	ctx := context.Background()
	res, err := esapi.IndexRequest{
		Index:      "tx",
		DocumentID: tx,
		Body:       strings.NewReader(string("{'foo': 'yes', 'bar': 'no'}")),
	}.Do(ctx, connections.GetEsClient())
	defer res.Body.Close()
	if err != nil {
		fmt.Println("Error indexing")
	}
	fmt.Println()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			fmt.Println("Error decoding index error message")
		}
		fmt.Printf("[%s] %s: %s\n", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
	}
}
