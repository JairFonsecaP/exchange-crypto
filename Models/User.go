package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
	Wallet   Wallet `json:"wallet"`
}

func (u *User) Valid(user *User) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password)) == nil
}

func (u *User) EncriptPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}
