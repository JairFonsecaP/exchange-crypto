package controllers

import (
	models "Coins/Models"
	"Coins/repositories"
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

		userStored, e := repositories.GetUserByUsername(&userInput)
		if e != nil {
			http.Error(w, "Username or password invalid", http.StatusUnauthorized)
			return
		}

		if userStored.Valid(&userInput) {
			token := services.GetToken(&userInput)
			_ = json.NewEncoder(w).Encode(&token)
		} else {
			http.Error(w, "Username or password invalid", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
