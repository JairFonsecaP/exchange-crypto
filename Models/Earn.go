package models

type Earn struct {
	Id    string `json:"id"`
	Coins []Coin `json:"coins"`
}

func newEarn() Earn {
	return Earn{
		Id:    "asdfsda",
		Coins: nil,
	}
}
