package server

import (
	"log"
	"net/http"

	"myprojet/rote"
)

func Server() {
	http.HandleFunc("/", rote.Artists)
	http.HandleFunc("/artist", rote.Artist)
	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
