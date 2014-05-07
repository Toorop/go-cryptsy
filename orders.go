package cryptsy

type Order struct {
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
	total    float64 `json:"total"`
}
