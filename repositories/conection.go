package repositories

import (
	"database/sql"
	"fmt"
)

const (
	db_user     = "exchange"
	db_password = "exchange"
	db_db       = "Exchange"
	db_addr     = "127.0.0.1"
)

func getConnection() (*sql.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", db_user, db_password, db_addr, db_db)
	db, err := sql.Open("mysql", s)
	return db, err
}
