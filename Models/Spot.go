package models

type Spot struct {
	Id    int64  `json:"id"`
	Coins []Coin `json:"coins"`
}

func newSpot() Spot {
	return Spot{
		Id:    1,
		Coins: nil,
	}
}
