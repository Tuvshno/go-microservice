package types

// PriceResponse is a struct to format the price JSON response
type PriceResponse struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}
