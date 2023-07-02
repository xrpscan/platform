package indexer

import (
	"encoding/json"
	"fmt"
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
	// If Fee is a concrete string, convert it to Int
	Fee, ok := tx["Fee"].(string)
	if ok {
		FeeInt, err := strconv.ParseInt(Fee, 10, 64)
		if err == nil {
			tx["Fee"] = FeeInt
		}
	}

	// If Amount is a concrete string, convert it to Int representing XRP Amount
	Amount, ok := tx["Amount"].(string)
	if ok {
		AmountInt, err := strconv.ParseInt(Amount, 10, 64)
		if err == nil {
			tx["Amount"] = AmountInt
		}
	}
	// If Amount is a IOU, rename it to AmountIOU
	AmountIOU, ok := tx["Amount"].(map[string]interface{})
	if ok {
		tx["AmountIOU"] = AmountIOU
		delete(tx, "Amount")
	}

	// If Amount is a concrete string, convert it to Int representing XRP Amount
	TakerGets, ok := tx["TakerGets"].(string)
	if ok {
		TakerGetsInt, err := strconv.ParseInt(TakerGets, 10, 64)
		if err == nil {
			tx["TakerGets"] = TakerGetsInt
		}
	}
	// If TakerGets is a IOU, rename it to TakerGetsIOU
	TakerGetsIOU, ok := tx["TakerGets"].(map[string]interface{})
	if ok {
		tx["TakerGetsIOU"] = TakerGetsIOU
		delete(tx, "TakerGets")
	}

	// If TakerPays is a concrete string, convert it to Int representing XRP Amount
	TakerPays, ok := tx["TakerPays"].(string)
	if ok {
		TakerPaysInt, err := strconv.ParseInt(TakerPays, 10, 64)
		if err == nil {
			tx["TakerPays"] = TakerPaysInt
		}
	}
	// If TakerPays is a IOU, rename it to TakerPaysIOU
	TakerPaysIOU, ok := tx["TakerPays"].(map[string]interface{})
	if ok {
		tx["TakerPaysIOU"] = TakerPaysIOU
		delete(tx, "TakerPays")
	}

	// If SendMax is a concrete string, convert it to Int representing XRP Amount
	SendMax, ok := tx["SendMax"].(string)
	if ok {
		SendMaxInt, err := strconv.ParseInt(SendMax, 10, 64)
		if err == nil {
			tx["SendMax"] = SendMaxInt
		}
	}
	// If SendMax is a IOU, rename it to SendMaxIOU
	SendMaxIOU, ok := tx["SendMax"].(map[string]interface{})
	if ok {
		tx["SendMaxIOU"] = SendMaxIOU
		delete(tx, "SendMax")
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

	// Fix tx.meta property
	meta, ok := tx["meta"].(map[string]interface{})
	if ok {
		delete(meta, "AffectedNodes")
		deliveredAmount, ok := meta["delivered_amount"].(string)
		if ok {
			deliveredAmountInt, err := strconv.ParseInt(deliveredAmount, 10, 64)
			if err == nil {
				meta["delivered_amount"] = deliveredAmountInt
			}
		}
		deliveredAmountIOU, ok := meta["delivered_amount"].(map[string]interface{})
		if ok {
			meta["delivered_amountIOU"] = deliveredAmountIOU
			delete(meta, "delivered_amount")
		}
		tx["meta"] = meta
	}
	return tx
}

func fixAmount(amount interface{}) (drops int64, iou map[string]interface{}, isIOU bool) {
	switch result := amount.(type) {
	case string:
		fmt.Println("Amount:", result)
		drops, err := strconv.ParseInt(result, 10, 64)
		if err == nil {
			return drops, nil, false
		}
	case map[string]interface{}:
		fmt.Println("Amount:", result)
		return 0, result, true
	}
	return 0, nil, false
}
