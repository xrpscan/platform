package indexer

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/platform/models"
	"github.com/xrpscan/xrpl-go"
	xrplgomodels "github.com/xrpscan/xrpl-go/models"
)

// Modify transacion object to normalize Amount-like and other fields.
func ModifyTransaction(tx map[string]interface{}) (map[string]interface{}, error) {
	// Detect network from NetworkID field. Default is XRPL Mainnet
	network := xrpl.NetworkXrplMainnet
	networkId, ok := tx["NetworkID"].(int)
	if ok {
		network = xrpl.GetNetwork(networkId)
	}

	// Rename tx.metaData property to tx.meta
	metaDataField := "metaData"
	_, ok2 := tx[metaDataField]
	if ok2 {
		tx["meta"] = tx[metaDataField]
		delete(tx, metaDataField)
	}

	// Add CTID field to the transaction if its missing
	_, hasCtid := tx["ctid"].(string)
	if !hasCtid {
		ctid, err := getCTID(tx["ledger_index"], tx["meta"], network)
		if err == nil {
			tx["ctid"] = ctid
		}
	}

	// Modify Fee from string to int64
	feeStr, ok := tx["Fee"].(string)
	if ok {
		fee, err := strconv.ParseInt(feeStr, 10, 64)
		if err == nil {
			tx["Fee"] = fee
		}
	}

	// Modify Amount-like fields listed in models.AmountFields
	for _, field := range models.AmountFields {
		modifyAmount(tx, field.String(), network)
	}

	// Modify contents of fields with Hex data
	for _, field := range models.HexFields {
		modifyHex(tx, field.String())
	}

	// Modify contents of fields with Dates
	for _, field := range models.DateFields {
		modifyDate(tx, field.String())
	}

	/*
	* Modify Memos field:
	* - Marshal Memos field to JSON []byte
	* - Unmarshal it to []models.Memos
	* - Mutate contents of individual Memo fields
	* - Marshal it back to map[string]interface{}
	* - Set the mutated value back to tx object
	 */
	if tx["Memos"] != nil {
		memosJSON, err := json.Marshal(tx["Memos"])
		if err == nil {
			var Memos []models.Memos
			err := json.Unmarshal(memosJSON, &Memos)
			if err == nil {
				for i := range Memos {
					Memos[i].Memo.MemoData = hexDecode(Memos[i].Memo.MemoData)
					Memos[i].Memo.MemoType = hexDecode(Memos[i].Memo.MemoType)
					Memos[i].Memo.MemoFormat = hexDecode(Memos[i].Memo.MemoFormat)
				}
			}
			tx["Memos"] = Memos
		}
	}

	// Modify Amount-like fields in meta
	meta, ok := tx["meta"].(map[string]interface{})
	if ok {
		// For simplicity, AffectedNodes field is dropped. This field may indexed
		// in a future release after due consideration.
		delete(meta, "AffectedNodes")
		modifyAmount(meta, models.DeliveredAmount.String(), network)
		modifyAmount(meta, models.Delivered_Amount.String(), network)
		tx["meta"] = meta
	}

	return tx, nil
}

func modifyAmount(tx map[string]interface{}, field string, network xrpl.Network) error {
	if tx[field] == nil {
		return nil
	}

	iou, ok := tx[field].(map[string]interface{})
	if ok {
		// TODO: Handle values expressed in scientific notation
		// snMatch, _ := regexp.MatchString(`^\d+[eE]\d+$`, iouValueStr)

		iouValueStr, ok := iou["value"].(string)
		if ok {
			iouValue, err := strconv.ParseFloat(iouValueStr, 64)
			if err != nil {
				logger.Log.Trace().Err(err).Str("field", field).Msg("IOU value error")
			} else {
				iou["value"] = iouValue
				tx[field] = iou
			}
		}
	}

	// If value is native asset, therefore represented as just a string, convert
	// it to Currency
	valueStr, ok := tx[field].(string)
	if ok {
		value, err := strconv.ParseInt(valueStr, 10, 64)
		if err == nil {
			tx[field] = map[string]interface{}{"currency": network.Asset(), "value": value}
		}
		tx[field] = map[string]interface{}{"currency": network.Asset(), "value": value}
	}
	return nil
}

func modifyHex(tx map[string]interface{}, field string) error {
	hex, ok := tx[field].(string)
	if ok {
		tx[field] = hexDecode(hex)
		fmt.Println("CTID:", tx["ctid"], "HEX:", hex, "==> ASCII:", tx[field])
	}
	return nil
}

func hexDecode(encoded string) string {
	decoded, err := hex.DecodeString(encoded)
	if err != nil {
		return encoded
	}
	return string(decoded)
}

func modifyDate(tx map[string]interface{}, field string) error {
	rippleTimestamp, ok := tx[field].(float64)
	if ok {
		newField := fmt.Sprintf("_%s", field)
		newTime := xrpl.RippleTimeToISOTime(int64(rippleTimestamp))
		// fmt.Println("TX:", tx["hash"], "Date:", newTime)
		tx[newField] = newTime
	}
	return nil
}

func getCTID(ledgerIndex interface{}, meta interface{}, networkId xrpl.Network) (string, error) {
	lgrIndex, ok := ledgerIndex.(float64)
	if !ok {
		return "", errors.New("cannot assert ledger_index as float64")
	}

	metaMSI, ok := meta.(map[string]interface{})
	if !ok {
		return "", errors.New("cannot parse meta field")
	}
	txnIndex, ok := metaMSI["TransactionIndex"].(float64)
	if !ok {
		return "", errors.New("cannot assert meta.TransactionIndex as float64")
	}

	ct := xrplgomodels.CTID{
		LedgerIndex:      uint64(lgrIndex),
		TransactionIndex: uint64(txnIndex),
		NetworkId:        uint64(networkId),
	}
	return ct.Encode(), nil
}
