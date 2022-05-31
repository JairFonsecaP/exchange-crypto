package repositories

import (
	models "Coins/Models"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func AddUser(u *models.User) {
	db, err := getConnection()
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	err = insertUser(db, u)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllUsers() (users []models.User, err error) {
	db, err := getConnection()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	res, e := db.Query("SELECT * FROM `Users`;")
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
	db, err := getConnection()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	q := "SELECT * FROM `Users` WHERE `Username` = ?;"
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
	q := "INSERT INTO `Users`(Name,Username,Password,Email) VALUES (?,?,?,?);"
	insert, err := db.Prepare(q)
	defer insert.Close()

	if err != nil {
		return err
	}
	_, err = insert.Exec(u.Name, u.Username, u.Password, u.Email)
	if err != nil {
		return err
	}
	return nil
}
