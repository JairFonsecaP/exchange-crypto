package controllers

import (
	models "Coins/Models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ListCoins(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&price_change_percentage=1h%2C7d%2C180d%2C1y")
	if err != nil {
		log.Fatal(err)
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	var coins []models.Coin

	err = json.Unmarshal(data, &coins)
	if err != nil {
		log.Fatal(err)
	}

	out, err := json.Marshal(coins)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func ListCoinsFiltered(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&price_change_percentage=1h%2C7d%2C180d%2C1y")
	if err != nil {
		log.Fatal(err)
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	var coins []models.Coin
	err = json.Unmarshal(data, &coins)
	if err != nil {
		log.Fatal(err)
	}
	filter := strings.ToLower(r.URL.Query().Get("search"))
	var coinsFiltered []models.Coin
	for _, v := range coins {
		if strings.Contains(strings.ToLower(v.Name), filter) || strings.Contains(strings.ToLower(v.Symbol), filter) {
			coinsFiltered = append(coinsFiltered, v)
		}
	}

	out, err := json.Marshal(coinsFiltered)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
