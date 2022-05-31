package models

type Wallet struct {
	Id         string `json:"id"`
	FiatPocket Fiat   `json:"fiatPocket"`
	SpotPocket Spot   `json:"spotPocket"`
	EarnPocket Earn   `json:"earnPocket"`
}

func newWallet() Wallet {
	w := Wallet{
		Id:         "asdf",
		FiatPocket: NewFiat(),
		SpotPocket: newSpot(),
		EarnPocket: newEarn(),
	}

	return w
}
