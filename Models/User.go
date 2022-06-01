package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
	Wallet   Wallet `json:"wallet"`
}

func (u *User) Valid(user *User) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password)) == nil
}

func (u *User) EncriptPassword() {
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hash)
}
