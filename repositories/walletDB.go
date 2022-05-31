package repositories

import (
	models "Coins/Models"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func createAWallet(w *models.Wallet) {
	db, err := getConnection()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = insertWallet(db, w)
	if err != nil {
		log.Fatal(err)
	}
}

func insertWallet(db *sql.DB, w *models.Wallet) error {
	q := "IN"
}
