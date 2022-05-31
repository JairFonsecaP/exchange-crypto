package models

const (
	AMOUNT_INITIAL = 50000.00
)

type Fiat struct {
	Id    string  `json:"id"`
	Total float64 `json:"total"`
}

func NewFiat() Fiat {
	return Fiat{
		Id:    "asdf", ///get hash
		Total: AMOUNT_INITIAL,
	}
}
