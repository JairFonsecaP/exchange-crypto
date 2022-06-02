package controllers

import (
	models "Coins/Models"
	"Coins/repositories"
	"Coins/services"
	"encoding/json"
	"io/ioutil"
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
			resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&price_change_percentage=1h%2C7d%2C180d%2C1y")
			if err != nil {
				http.Error(w, "Something went error", http.StatusNotFound)
				return
			}

			data, _ := ioutil.ReadAll(resp.Body)
			var coins []models.Coin

			err = json.Unmarshal(data, &coins)
			e = repositories.GetSpotCoinsActive(&wallet, coins)
			if e != nil {
				http.Error(w, "Something went error", http.StatusNotFound)
				return
			}
			e = repositories.GetEarnCoinsActive(&wallet, coins)
			if e != nil {
				http.Error(w, "Something went error", http.StatusNotFound)
				return
			}
			e = services.CompleteData(&wallet.SpotPocket, &wallet.EarnPocket)
			if e != nil {
				http.Error(w, "Something went error", http.StatusNotFound)
				return
			}
			out, err := json.MarshalIndent(&wallet, "", "    ")
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
