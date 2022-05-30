package auth

import (
	"Coins/services"
	"net/http"
)

func ValidateToken(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if services.ValidateToken(token) {
			next.ServeHTTP(w, r)
			return
		}
		//redirigir a login.html
	}
}
