package cryptsy

type Trade struct {
	Id       int64   `json:"id"`
	Time     string  `json:"time"`
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
	total    float64 `json:"total"`
}
