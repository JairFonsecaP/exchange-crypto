package controllers

import (
	models "Coins/Models"
	"Coins/repositories"
	"Coins/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Error", 501)
			return
		}
		var user models.User
		err = json.Unmarshal(data, &user)
		if err != nil {
			http.Error(w, "User exists please use other username", http.StatusNotAcceptable)
			return
		}
		userTemp, er := repositories.GetUserByUsername(&user)
		if er != nil || userTemp.Id != 0 {
			http.Error(w, "User exists please use other username", http.StatusNotAcceptable)
			return
		}

		if err != nil {
			http.Error(w, "Error", 501)
			return
		}
		user.EncriptPassword()
		err = repositories.AddUser(&user)
		if err != nil {
			http.Error(w, "Error", 501)
			return
		}
		err = createWallet(&user)
		if err != nil {
			http.Error(w, "Error", 501)
			return
		}

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Username or password invalid", http.StatusUnauthorized)
			return
		}
		var userInput models.User
		err = json.Unmarshal(data, &userInput)

		if err != nil {
			http.Error(w, "Username or password invalid", http.StatusUnauthorized)
			return
		}

		userStored, e := repositories.GetUserByUsername(&userInput)
		if e != nil {
			http.Error(w, "Username or password invalid", http.StatusUnauthorized)
			return
		}

		if userStored.Valid(&userInput) {
			token := services.GetToken(&userInput)
			cookie := &http.Cookie{
				Name:     "token",
				Value:    token.Token,
				Path:     "/",
				SameSite: http.SameSiteNoneMode,
				Secure:   true,
			}
			http.SetCookie(w, cookie)
			out, err := json.MarshalIndent(token, "", "    ")
			if err != nil {
				http.Error(w, "Something went error", http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(out)
		} else {
			http.Error(w, "Username or password invalid", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
