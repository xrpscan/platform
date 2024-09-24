package indexer

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

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
			tx["_Fee"] = float64(fee) / models.DROPS_IN_XRP // Drops to XRP units
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

// Field legend
//
// value: Original value of the value field in string
//
// _value: Normalized value of the field in float64. XRP values are converted
// from Drops to XRP units. IOU value fields are not changed.
//
// currency: Original value of the currency field in ISO-4217 string or 160-bit HEX
//
// _currency: ISO-4217 currency code or HEX decoded currency code
func modifyAmount(tx map[string]interface{}, field string, network xrpl.Network) error {
	if tx[field] == nil {
		return nil
	}

	// If value is an IOU with currency, issuer and value fields
	iou, ok := tx[field].(map[string]interface{})
	if ok {
		// TODO: Handle values expressed in scientific notation
		// snMatch, _ := regexp.MatchString(`^\d+[eE]\d+$`, iouValueStr)

		// Modify IOU's value fields
		iouValueStr, ok := iou["value"].(string)
		if ok {
			iouValue, err := strconv.ParseFloat(iouValueStr, 64)
			if err != nil {
				logger.Log.Trace().Err(err).Str("field", field).Msg("IOU value error")
			} else {
				iou["_value"] = iouValue
			}
		}

		// Modify IOU's currency field
		currency, okc := iou["currency"].(string)
		if okc {
			iou["_currency"] = getCurrencyCode(currency)
		}

		// Finally, set the iou field back
		tx[field] = iou
	}

	// If value is native asset, represented as a string value, convert it to
	// Currency object with currency and value fields. Issuer field is not set.
	valueStr, ok := tx[field].(string)
	if ok {
		value, err := strconv.ParseInt(valueStr, 10, 64)
		if err == nil {
			tx[field] = map[string]interface{}{
				"value":     valueStr,
				"_value":    float64(value) / models.DROPS_IN_XRP,
				"currency":  network.Asset(),
				"_currency": network.Asset(),
				"native":    true,
			}
		} else {
			tx[field] = map[string]interface{}{
				"value":     valueStr,
				"currency":  network.Asset(),
				"_currency": network.Asset(),
				"native":    true,
			}
			logger.Log.Error().Err(err).Str("field", field).Str("value", valueStr).Msg("Native value error")
		}
	}
	return nil
}

func modifyHex(tx map[string]interface{}, field string) error {
	hex, ok := tx[field].(string)
	if ok {
		tx[field] = hexDecode(hex)
	}
	return nil
}

func hexDecode(encoded string) string {
	nullChar, _ := hex.DecodeString("00")
	decoded, err := hex.DecodeString(encoded)
	if err != nil {
		return encoded
	}
	return strings.TrimRight(string(decoded), string(nullChar))
}

// Returns human readable currency code for XRP, 3-letter and HEX encoded currency codes
func getCurrencyCode(currency string) string {
	if len(currency) == 40 {
		return hexDecode(currency)
	}
	return currency
}

func modifyDate(tx map[string]interface{}, field string) error {
	rippleTimestamp, ok := tx[field].(float64)
	if ok {
		newField := fmt.Sprintf("_%s", field)
		newTime := xrpl.RippleTimeToISOTime(int64(rippleTimestamp))
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
