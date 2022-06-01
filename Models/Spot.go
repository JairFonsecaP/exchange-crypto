package models

type Spot struct {
	Id    int64  `json:"id"`
	Coins []Coin `json:"coins"`
}

func createPoket(id int64) Spot {
	return Spot{
		Id:    id,
		Coins: nil,
	}
}
