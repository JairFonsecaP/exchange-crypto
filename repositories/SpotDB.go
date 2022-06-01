package repositories

import (
	models "Coins/Models"
	"database/sql"
)

func insertSpot(db *sql.DB, w *models.Wallet) error {
	q := "INSERT INTO `Spot`(id) VALUES(?);"
	insert, err := db.Prepare(q)
	defer insert.Close()
	if err != nil {
		return err
	}
	res, e := insert.Exec(w.SpotPocket.Id)
	if e != nil {
		return e
	}
	i, er := res.LastInsertId()
	if er != nil {
		return er
	}
	w.SpotPocket.Id = i
	return nil
}
