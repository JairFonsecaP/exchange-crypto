package services

import (
	models "Coins/Models"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const singKey = "Isi-Montreal,programming"

type JWT struct {
	Token string `json:"token"`
}

func GetToken(user *models.User) *JWT {
	token := sign(user)
	return &JWT{Token: token}
}

func sign(user *models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.Id,
		"name": user.Name,
		"exp":  time.Now().Add(60 * time.Minute),
	})
	mySigningKey := []byte(singKey)
	t, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return t
}

func ValidateToken(s string) bool {
	mySigningKey := []byte(singKey)
	t, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return t.Valid

}
