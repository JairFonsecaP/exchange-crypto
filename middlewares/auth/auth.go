package auth

import (
	"Coins/services"
	"net/http"
)

func ValidateToken(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := r.Cookie("token")
		if services.ValidateToken(t.Value) {
			next.ServeHTTP(w, r)
			return
		}
	}
}
