package services

import (
	models "Coins/Models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CompleteData(s *models.Spot, e *models.Earn) error {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&price_change_percentage=1h%2C7d%2C180d%2C1y")
	if err != nil {
		log.Fatal(err)
		return err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	var coins []models.Coin

	err = json.Unmarshal(data, &coins)

	for _, coin := range coins {
		for _, c := range s.Coins {
			if coin.Id == c.Id {
				c.Name = coin.Name
				c.Image = coin.Image
				c.Symbol = coin.Symbol
				c.Price = coin.Price
				c.Price_change_percentage_1h_in_currency = coin.Price_change_percentage_1h_in_currency
				c.Price_change_percentage_24h = coin.Price_change_percentage_24h
				c.Price_change_percentage_7d_in_currency = coin.Price_change_percentage_7d_in_currency
				c.Price_change_percentage_1y_in_currency = coin.Price_change_percentage_1y_in_currency
				break
			}
		}
		for _, ce := range e.Coins {
			if coin.Id == ce.Id {
				ce.Name = coin.Name
				ce.Image = coin.Image
				ce.Symbol = coin.Symbol
				ce.Price = coin.Price
				ce.Price_change_percentage_1h_in_currency = coin.Price_change_percentage_1h_in_currency
				ce.Price_change_percentage_24h = coin.Price_change_percentage_24h
				ce.Price_change_percentage_7d_in_currency = coin.Price_change_percentage_7d_in_currency
				ce.Price_change_percentage_1y_in_currency = coin.Price_change_percentage_1y_in_currency
				break
			}
		}
	}
	return nil
}
