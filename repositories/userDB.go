package repositories

import (
	models "Coins/Models"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func AddUser(u *models.User) error {
	db, err := getConnection()
	defer db.Close()

	if err != nil {
		return err
	}
	err = insertUser(db, u)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers() (users []models.User, err error) {
	db, err := getConnection()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	res, e := db.Query("SELECT * FROM `User`;")
	if e != nil {
		return users, err
	}
	for res.Next() {
		var user models.User
		err = res.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.Email)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByUsername(user *models.User) (u models.User, err error) {
	db, er := getConnection()
	defer db.Close()
	if er != nil {
		log.Fatal(er)
		return u, er
	}
	q := "SELECT * FROM `User` WHERE `Username` = ?;"
	res, e := db.Query(q, user.Username)
	if e != nil {
		return u, e
	}
	for res.Next() {
		err = res.Scan(&u.Id, &u.Name, &u.Username, &u.Password, &u.Email)
		if err != nil {
			return u, err
		}
	}
	return u, nil
}

func insertUser(db *sql.DB, u *models.User) error {
	q := "INSERT INTO `User`(`Name`,`Username`,`Password`,`Email`) VALUES (?,?,?,?);"
	insert, err := db.Prepare(q)
	defer insert.Close()

	if err != nil {
		log.Fatal(err)
		return err
	}
	res, err := insert.Exec(u.Name, u.Username, u.Password, u.Email)
	if err != nil {
		log.Fatal(err)
		return err
	}
	u.Id, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
