package models

// Amount-like fields
type AmountField int8

const (
	Amount AmountField = iota
	DeliverMax
	DeliveredAmount
	Delivered_Amount
	DeliverMin
	SendMax
	LimitAmount
	TakerGets
	TakerPays
	Currency_
	NFTokenBrokerFee
	Amount2
	Asset
	Asset2
	BidMin
	BidMax
	EPrice
	LPTokenIn
	LPTokenOut
)

// Converts AmountField to its string representation
func (af AmountField) String() string {
	switch af {
	case Amount:
		return "Amount"
	case DeliverMax:
		return "DeliverMax"
	case DeliveredAmount:
		return "DeliveredAmount"
	case Delivered_Amount:
		return "delivered_amount"
	case DeliverMin:
		return "DeliverMin"
	case SendMax:
		return "SendMax"
	case LimitAmount:
		return "LimitAmount"
	case TakerGets:
		return "TakerGets"
	case TakerPays:
		return "TakerPays"
	case Currency_:
		return "currency"
	case NFTokenBrokerFee:
		return "NFTokenBrokerFee"
	case Amount2:
		return "Amount2"
	case Asset:
		return "Asset"
	case Asset2:
		return "Asset2"
	case BidMin:
		return "BidMin"
	case BidMax:
		return "BidMax"
	case EPrice:
		return "EPrice"
	case LPTokenIn:
		return "LPTokenIn"
	case LPTokenOut:
		return "LPTokenOut"
	}
	return ""
}

// A slice of all known Amount-like fields in XRPL transaction.
var AmountFields = []AmountField{
	Amount,
	DeliverMax,
	DeliveredAmount,
	Delivered_Amount,
	DeliverMin,
	SendMax,
	LimitAmount,
	TakerGets,
	TakerPays,
	Currency_,
	NFTokenBrokerFee,
	Amount2,
	Asset,
	Asset2,
	BidMin,
	BidMax,
	EPrice,
	LPTokenIn,
	LPTokenOut,
}

// Fields with hex encoded data
type HexField int8

const (
	Domain HexField = iota
	Data
	URI
	Provider
	AssetClass
)

// Convert HexField to its string representation
func (hf HexField) String() string {
	switch hf {
	case Domain:
		return "Domain"
	case Data:
		return "Data"
	case URI:
		return "URI"
	case Provider:
		return "Provider"
	case AssetClass:
		return "AssetClass"
	}
	return ""
}

// Slice of all known hex fields in XRPL transaction
var HexFields = []HexField{
	Domain,
	Data,
	URI,
	Provider,
	AssetClass,
}

// Fields with Date
type DateField int8

const (
	Date DateField = iota
	Expiration
	CancelAfter
	FinishAfter
	LastUpdateTime
)

// Convert Date field to its string representation as present in tx object
func (df DateField) String() string {
	switch df {
	case Date:
		return "date"
	case Expiration:
		return "Expiration"
	case CancelAfter:
		return "CancelAfter"
	case FinishAfter:
		return "FinishAfter"
	case LastUpdateTime:
		return "LastUpdateTime"
	}
	return ""
}

// A slice of all known Date fields in XRPL transaction
var DateFields = []DateField{
	Date,
	Expiration,
	CancelAfter,
	FinishAfter,
	LastUpdateTime,
}
