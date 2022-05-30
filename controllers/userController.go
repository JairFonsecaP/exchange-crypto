package controllers

import (
	models "Coins/Models"
	"Coins/services"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Fatal(err)
			return
		}
		var userInput models.User
		err = json.Unmarshal(data, &userInput)

		if err != nil {
			log.Fatal(err)
			return
		}
		//poner la logica para traer de db
		var user models.User
		user = models.User{Name: "jair", Username: "jair", Password: "1234"}

		if user.Valid(&userInput) {
			token := services.GetToken(&user)
			_ = json.NewEncoder(w).Encode(&token)
		} else {
			http.Error(w, "Username or password invalid", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
