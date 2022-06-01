package models

type Earn struct {
	Id    int64  `json:"id"`
	Coins []Coin `json:"coins"`
}
