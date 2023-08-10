package models

type Meta struct {
	// AffectedNodes // Unimplemented
	DeliveredAmount   Currency `json:"delivered_amount,omitempty"`
	Delivered_Amount  Currency `json:"DeliveredAmount,omitempty"`
	TransactionIndex  int      `json:"TransactionIndex,omitempty"`
	TransactionResult string   `json:"TransactionResult,omitempty"`
}
