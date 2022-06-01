package repositories

import (
	models "Coins/Models"
	"database/sql"
	"log"
)

func CreateAWallet(w *models.Wallet) error {
	db, err := getConnection()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = createWallet(db, w)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func createWallet(db *sql.DB, w *models.Wallet) error {
	err := insertSpot(db, w)
	if err != nil {
		return err
	}
	err = insertFiat(db, w)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = insertEarn(db, w)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = insertWallet(db, w)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func insertWallet(db *sql.DB, w *models.Wallet) error {
	q := "INSERT INTO `wallet`(Spot_Pocket,Earn_Pocket,Fiat_Pocket,Id_User) VALUES(?,?,?,?);"
	insert, err := db.Prepare(q)
	defer insert.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	res, e := insert.Exec(w.SpotPocket.Id, w.EarnPocket.Id, w.FiatPocket.Id, w.Id_User)
	if e != nil {
		log.Fatal(e)
		return e
	}
	i, er := res.LastInsertId()

	if er != nil {
		log.Fatal(er)
		return er
	}
	w.Id = i
	return nil
}
