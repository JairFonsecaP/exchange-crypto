package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "static/html/index.html")
}

func LoginView(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "static/html/login.html")
}

func SignUpView(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "static/html/signup.html")
}

func DashboardView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "secure/html/dashboard.html")
	}
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}

}
