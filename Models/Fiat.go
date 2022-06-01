package models

const (
	AMOUNT_INITIAL = 50000.00
)

type Fiat struct {
	Id    int64   `json:"id"`
	Total float64 `json:"total"`
}

func NewFiat() Fiat {
	return Fiat{Id: -1, Total: AMOUNT_INITIAL}
}
