package main

import (
	"log"
	"net/http"

	"http_svr/server"
)

func main() {

	svr := server.NewServer()
	http.HandleFunc("/", svr.ServeHTTP)

	log.Print("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
