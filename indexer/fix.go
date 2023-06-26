package indexer

import (
	"encoding/json"
	"strconv"

	"github.com/xrpscan/platform/logger"
)

func FixTransactionObject(transaction []byte) ([]byte, error) {
	// Unmarshal transaction JSON and handle it as map[string]interface{}
	var tx map[string]interface{}
	if err := json.Unmarshal(transaction, &tx); err != nil {
		logger.Log.Error().Err(err).Msg("JSON Unmarshal error in fix tx")
		return transaction, err
	}

	txAmount := FixTransactionAmount(tx)
	txMetaData := FixTransactionMetaData(txAmount)

	// Marshal transaction object back to JSON
	result, err := json.Marshal(txMetaData)
	if err != nil {
		return transaction, err
	}

	return result, nil
}

func FixTransactionAmount(tx map[string]interface{}) map[string]interface{} {
	Fee, ok := tx["Fee"].(string)
	if ok {
		FeeInt, err := strconv.ParseInt(Fee, 10, 64)
		if err == nil {
			tx["Fee"] = FeeInt
		}
	}

	Amount, ok := tx["Amount"].(string)
	if ok {
		AmountInt, err := strconv.ParseInt(Amount, 10, 64)
		if err == nil {
			tx["Amount"] = AmountInt
		}
	}

	TakerGets, ok := tx["TakerGets"].(string)
	if ok {
		TakerGetsInt, err := strconv.ParseInt(TakerGets, 10, 64)
		if err == nil {
			tx["TakerGets"] = TakerGetsInt
		}
	}

	TakerPays, ok := tx["TakerPays"].(string)
	if ok {
		TakerPaysInt, err := strconv.ParseInt(TakerPays, 10, 64)
		if err == nil {
			tx["TakerPays"] = TakerPaysInt
		}
	}

	SendMax, ok := tx["SendMax"].(string)
	if ok {
		SendMaxInt, err := strconv.ParseInt(SendMax, 10, 64)
		if err == nil {
			tx["SendMax"] = SendMaxInt
		}
	}
	return tx
}

func FixTransactionMetaData(tx map[string]interface{}) map[string]interface{} {
	// Rename tx.metaData property to tx.meta
	_, ok := tx["metaData"]
	if ok {
		tx["meta"] = tx["metaData"]
		delete(tx, "metaData")
	}

	// Drop tx.meta.AffectedNodes property entirely
	meta, ok := tx["meta"].(map[string]interface{})
	if ok {
		delete(meta, "AffectedNodes")
		tx["meta"] = meta
	}
	return tx
}
