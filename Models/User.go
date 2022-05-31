package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"` ///corregir
	// Email string `json:"email"`
}

func (u *User) Valid(userInput *User) bool {
	return u.Username == userInput.Username && bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userInput.Password)) == nil
}

func (u *User) EncriptPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}
