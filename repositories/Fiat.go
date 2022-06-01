package repositories

import (
	models "Coins/Models"
	"database/sql"
	"log"
)

func insertFiat(db *sql.DB, w *models.Wallet) error {

	q := "INSERT INTO `Fiat`(total) VALUES(?);"
	insert, err := db.Prepare(q)
	defer insert.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	res, e := insert.Exec(w.FiatPocket.Total)
	if e != nil {
		log.Fatal(e)
		return e
	}
	i, er := res.LastInsertId()
	if er != nil {
		log.Fatal(er)
		return er
	}
	w.FiatPocket.Id = i
	return nil

}
