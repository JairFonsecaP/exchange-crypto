package main

import (
	"Coins/router"
	"log"
	"net/http"
)

type webServer struct {
	server *http.Server
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	server := newWebServer("8080")
	go router.Router()
	server.start()
}

func newWebServer(port string) *webServer {
	server := &http.Server{
		Addr: ":" + port,
	}

	return &webServer{
		server: server,
	}
}

func (w *webServer) start() {
	log.Println("server start at port" + w.server.Addr)
	log.Fatal(w.server.ListenAndServe())
}
