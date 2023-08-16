package indexer

import (
	"encoding/json"

	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/platform/models"
)

// Raw transaction object represented as a map-string-interface
type MapStringInterface map[string]interface{}

// Modify transacion object to normalize Amount-like and other fields.
func ModifyTransaction(transaction []byte) ([]byte, error) {
	// Unmarshal transaction and handle it as map[string]interface{}
	var tx map[string]interface{}

	if err := json.Unmarshal(transaction, &tx); err != nil {
		logger.Log.Error().Err(err).Msg("JSON Unmarshal error in ModifyTransaction")
		return transaction, err
	}

	// Modify Amount-like fields listed in models.AmountFields
	for _, field := range models.AmountFields {
		ModifyAmount(tx, field.String())
	}

	// Rename tx.metaData property to tx.meta
	metaDataField := "metaData"
	_, ok := tx[metaDataField]
	if ok {
		tx["meta"] = tx[metaDataField]
		delete(tx, metaDataField)
	}

	// Modify Amount-like fields in meta
	meta, ok := tx["meta"].(map[string]interface{})
	if ok {
		// For simplicity, AffectedNodes field is dropped. This field may indexed
		// in a future release after due consideration.
		delete(meta, "AffectedNodes")
		ModifyAmount(meta, models.DeliveredAmount.String())
		ModifyAmount(meta, models.Delivered_Amount.String())
		tx["meta"] = meta
	}

	// Marshal transaction object back to []byte
	result, err := json.Marshal(tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ModifyAmount(tx MapStringInterface, field string) error {
	value, ok := tx[field].(string)
	if ok {
		tx[field] = MapStringInterface{"currency": models.XRP, "value": value}
	}
	return nil
}
