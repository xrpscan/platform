package models

const XRP = "XRP"

type Currency struct {
	Currency string `json:"currency,omitempty"`
	Issuer   string `json:"issuer,omitempty"`
	Value    int64  `json:"value,omitempty"`
	Native   bool   `json:"native,omitempty"`
}

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
}
