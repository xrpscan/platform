package producers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/platform/models"
	"github.com/xrpscan/xrpl-go"
)

/*
* Verifies structure of `ledger` websocket command, iterates over `transactions`
* slice, and calls ProduceTransaction for each transaction object.
*
* These transaction objects have slightly different structure, as compared to
* transactions returned via `tx` and `account_tx` commands.
 */
func ProduceTransactions(w *kafka.Writer, message []byte) {
	// var res xrpl.BaseResponse
	// if err := json.Unmarshal(message, &res); err != nil {
	// 	logger.Log.Error().Err(err).Msg("JSON Unmarshal error")
	// 	return
	// }
	// ledgerIndex := strconv.Itoa(int(res["ledger_index"].(float64)))

	var ledger models.LedgerStream
	if err := json.Unmarshal(message, &ledger); err != nil {
		logger.Log.Error().Err(err).Msg("JSON Unmarshal error")
		return
	}

	// Type check ledger object in message and call Ledger.FetchTransactions on it

	// Fetch all transactions included in this ledger from XRPL server
	txResponse, err := ledger.FetchTransactions()
	if err != nil {
		logger.Log.Error().Uint32("ledger_index", ledger.LedgerIndex).Err(err).Msg(err.Error())
		return
	}

	// Verify if result.ledger.transactions property is present
	txResult, ok := txResponse["result"].(map[string]interface{})
	if !ok {
		logger.Log.Error().Uint32("ledger_index", ledger.LedgerIndex).Msg("Tx response has no result property")
		return
	}
	txLedger, ok := txResult["ledger"].(map[string]interface{})
	if !ok {
		logger.Log.Error().Uint32("ledger_index", ledger.LedgerIndex).Msg("Tx response has no result.ledger property")
		return
	}
	txs, ok := txLedger["transactions"].([]interface{})
	if !ok {
		logger.Log.Error().Uint32("ledger_index", ledger.LedgerIndex).Msg("Tx response has no result.ledger.transactions property")
		return
	}

	// Iterate over transactions slice and submit each transaction to Kafka topic
	for _, txo := range txs {
		tx, ok := txo.(map[string]interface{})
		if !ok {
			logger.Log.Error().Uint32("ledger_index", ledger.LedgerIndex).Err(err).Msg("Error asserting transaction type")
			return
		}

		// Transactions fetched by `ledger` command do not have date, validated and
		// ledger_index fields. Populating these tx fields from ledger data.
		tx["ledger_index"] = ledger.LedgerIndex
		tx["date"] = ledger.LedgerTime
		tx["validated"] = true

		txJSON, err := json.Marshal(tx)
		if err != nil {
			logger.Log.Error().Uint32("ledger_index", ledger.LedgerIndex).Err(err).Msg("Error Marshaling transaction")
			return
		}
		ProduceTransaction(w, txJSON)
	}
}

/*
* Submits transaction object to Kafka
 */
func ProduceTransaction(w *kafka.Writer, message []byte) {
	var res xrpl.BaseResponse
	if err := json.Unmarshal(message, &res); err != nil {
		return
	}

	messageKey, ok := res["hash"].(string)
	if !ok {
		logger.Log.Error().Msg("Tx object has no hash, aborting.")
		return
	}

	err := w.WriteMessages(context.Background(),
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
