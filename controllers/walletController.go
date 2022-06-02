package controllers

import (
	models "Coins/Models"
	"Coins/repositories"
	"encoding/json"
	"net/http"
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
func GetWalletByUserId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			wallet, e := repositories.GetWalletByUserId(&models.User{Id: 1})
			if e != nil {
				http.Error(w, "Something went error", http.StatusNotFound)
				return
			}
			out, err := json.MarshalIndent(wallet, "", "    ")
			if err != nil {
				http.Error(w, "Something went error", http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(out)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
