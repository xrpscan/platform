package producers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/xrpl-go"
)

func FetchTransaction(ledgerIndex string) (xrpl.BaseResponse, error) {
	requestId := fmt.Sprintf("ledger.%s.tx", ledgerIndex)
	request := xrpl.BaseRequest{
		"id":           requestId,
		"command":      "ledger",
		"ledger_index": ledgerIndex,
		"transactions": true,
		"expand":       true,
	}
	response, err := connections.XrplClient.Request(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func ProduceTransaction(w *kafka.Writer, message []byte) {
	var res xrpl.BaseResponse
	if err := json.Unmarshal(message, &res); err != nil {
		return
	}
	ledgerIndex := strconv.Itoa(int(res["ledger_index"].(float64)))

	txResponse, err := FetchTransaction(ledgerIndex)
	if err != nil {
		logger.Log.Error().Str("ledger_index", ledgerIndex).Msg(err.Error())
	}
	txResult, ok := txResponse["result"].(map[string]interface{})
	if !ok {
		logger.Log.Error().Str("ledger_index", ledgerIndex).Msg("Tx response has no result property")
		return
	}
	txLedger, ok := txResult["ledger"].(map[string]interface{})
	if !ok {
		logger.Log.Error().Str("ledger_index", ledgerIndex).Msg("Tx response has no result.ledger property")
		return
	}
	txs, ok := txLedger["transactions"].([]interface{})
	if !ok {
		logger.Log.Error().Str("ledger_index", ledgerIndex).Msg("Tx response has no result.ledger.transactions property")
		return
	}

	for _, tx := range txs {
		txObject, ok := tx.(map[string]interface{})
		if !ok {
			logger.Log.Error().Str("ledger_index", ledgerIndex).Msg("Tx object refused to typecast")
			return
		}
		messageKey, ok := txObject["hash"].(string)
		if !ok {
			logger.Log.Error().Str("ledger_index", ledgerIndex).Msg("Tx object has no hash")
			return
		}
		message, err := json.Marshal(txObject)
		if err != nil {
			log.Printf("Failed to Marshal tx object: %s", err)
			return
		}

		// Write to Kafka topic config.TopicTx()
		err = w.WriteMessages(context.Background(),
			kafka.Message{
				Topic: config.TopicTransactions(),
				Key:   []byte(messageKey),
				Value: message,
			},
		)
		if err != nil {
			log.Printf("Failed to write message: %s", err)
		}
	}
}
