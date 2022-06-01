package controllers

import (
	models "Coins/Models"
	"Coins/repositories"
)

func createWallet(u *models.User) error {
	wallet := models.NewWallet()
	wallet.Id_User = u.Id
	e := repositories.CreateAWallet(wallet)
	if e != nil {
		return e
	}
	return nil
}
