package router

import (
	"Coins/controllers"
	"Coins/middlewares/auth"
	"net/http"
)

func Router() {
	//Router Wallets
	http.HandleFunc("/api/wallet/get", auth.ValidateToken(controllers.GetWalletByUserId()))

	//Router Coins
	http.HandleFunc("/api/coins/list", controllers.ListCoins)
	http.HandleFunc("/api/coins/filter", controllers.ListCoinsFiltered)

	//Router Users
	http.HandleFunc("/api/user/signup", controllers.SignUp)
	http.HandleFunc("/api/user/login", controllers.Login)

	//Router Views
	http.HandleFunc("/signup", controllers.SignUpView)
	http.HandleFunc("/login", controllers.LoginView)
	http.HandleFunc("/dashboard", auth.ValidateToken(controllers.DashboardView()))
	http.HandleFunc("/", controllers.Index)
}
