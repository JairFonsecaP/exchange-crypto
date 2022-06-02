package models

import "time"

type Coin struct {
	Id                                     string    `json:"id"`
	Name                                   string    `json:"name"`
	Image                                  string    `json:"image"`
	Symbol                                 string    `json:"symbol"`
	Price                                  float64   `json:"current_price"`
	Price_change_percentage_1h_in_currency float64   `json:"price_change_percentage_1h_in_currency"`
	Price_change_percentage_24h            float64   `json:"price_change_percentage_24h"`
	Price_change_percentage_7d_in_currency float64   `json:"price_change_percentage_7d_in_currency"`
	Price_change_percentage_1y_in_currency float64   `json:"price_change_percentage_1y_in_currency"`
	BuyPrice                               float64   `json:"buyprice,omitempty"`
	Quantity                               float64   `json:"quantity,omitempty"`
	SellPrice                              float64   `json:"sellprice,omitempty"`
	BuyDate                                time.Time `json:"buy_date,omitempty"`
	SellDate                               time.Time `json:"sell_date,omitempty"`
	// Market_cap                       float64 `json:"market_cap"`
	// Market_cap_rank                  float64 `json:"market_cap_rank"`
	// Fully_diluted_valuation          float64 `json:"fully_diluted_valuation"`
	// Total_volume                     float64 `json:"Total_volume"`
	// High_24h                         float64 `json:"high_24h"`
	// Low_24h                          float64 `json:"low_24h"`
	// Price_change_24h                 float64 `json:"price_change_24h"`
	// Market_cap_change_24h            float64 `json:"market_cap_change_24h"`
	// Market_cap_change_percentage_24h float64 `json:"market_cap_change_percentage_24h"`
	// Circulating_supply               float64 `json:"circulating_supply"`
	// Total_supply                     float64 `json:"total_supply"`
	// Max_supply                       float64 `json:"max_supply"`
	// Ath                              float64 `json:"ath"`
}
