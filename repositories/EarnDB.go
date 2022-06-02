package repositories

import (
	models "Coins/Models"
	"database/sql"
	"log"
)

func insertEarn(db *sql.DB, w *models.Wallet) error {
	q := "INSERT INTO `Earn`(id) VALUES(?);"
	insert, err := db.Prepare(q)
	defer insert.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	res, e := insert.Exec(w.EarnPocket.Id)
	if e != nil {
		log.Fatal(e)
		return e
	}
	i, er := res.LastInsertId()
	if er != nil {
		log.Fatal(er)
		return er
	}
	w.EarnPocket.Id = i
	return nil
}

func GetEarnCoinsActive(w *models.Wallet, coins []models.Coin) (err error) {
	db, er := getConnection()
	defer db.Close()
	if er != nil {
		log.Fatal(er)
		return er
	}
	q := "SELECT id_coin, buy_price, buy_date " +
		"from earncoins " +
		"where earn_pocket = ? and sell_price is null and sell_date is null;"
	res, e := db.Query(q, w.EarnPocket.Id)
	if e != nil {
		return e
	}

	for res.Next() {
		var c models.Coin
		err = res.Scan(&c.Id, &c.BuyPrice, &c.BuyDate)
		if err != nil {
			log.Fatal(err)
			return err
		}
		for _, coin := range coins {
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
		w.EarnPocket.Coins = append(w.EarnPocket.Coins, c)
	}
	return nil
}
