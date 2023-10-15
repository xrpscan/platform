package indexer

import (
	"encoding/json"

	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/platform/models"
	"github.com/xrpscan/xrpl-go"
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

	// Detect network from NetworkID field. Default is XRPL Mainnet
	network := xrpl.NetworkXrplMainnet
	networkId, ok := tx["NetworkID"].(int)
	if ok {
		network = xrpl.GetNetwork(networkId)
	}

	// Modify Amount-like fields listed in models.AmountFields
	for _, field := range models.AmountFields {
		ModifyAmount(tx, field.String(), network)
	}

	// Rename tx.metaData property to tx.meta
	metaDataField := "metaData"
	_, ok2 := tx[metaDataField]
	if ok2 {
		tx["meta"] = tx[metaDataField]
		delete(tx, metaDataField)
	}

	// Modify Amount-like fields in meta
	meta, ok := tx["meta"].(map[string]interface{})
	if ok {
		// For simplicity, AffectedNodes field is dropped. This field may indexed
		// in a future release after due consideration.
		delete(meta, "AffectedNodes")
		ModifyAmount(meta, models.DeliveredAmount.String(), network)
		ModifyAmount(meta, models.Delivered_Amount.String(), network)
		tx["meta"] = meta
	}

	// Marshal transaction object back to []byte
	result, err := json.Marshal(tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ModifyAmount(tx MapStringInterface, field string, network xrpl.Network) error {
	value, ok := tx[field].(string)
	if ok {
		tx[field] = MapStringInterface{"currency": network.Asset(), "value": value}
	}
	return nil
}
