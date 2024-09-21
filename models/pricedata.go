package models

type PriceData struct {
	BaseAsset  string `json:"BaseAsset,omitempty"`
	QuoteAsset string `json:"QuoteAsset,omitempty"`
	AssetPrice uint64 `json:"AssetPrice,omitempty"`
	Scale      uint8  `json:"Scale,omitempty"`
}
