package models

type Wallet struct {
	Id         int64 `json:"id"`
	FiatPocket Fiat  `json:"fiatPocket"`
	SpotPocket Spot  `json:"spotPocket"`
	EarnPocket Earn  `json:"earnPocket"`
	Id_User    int64 `json:"idUser"`
}

func NewWallet() *Wallet {
	var spot Spot
	var earn Earn
	fiat := NewFiat()

	wallet := Wallet{
		FiatPocket: fiat,
		SpotPocket: spot,
		EarnPocket: earn,
	}
	return &wallet
}
