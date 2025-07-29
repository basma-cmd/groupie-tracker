package rote

import (
	"html/template"
	"net/http"
)

func Artist(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "500 error internal server", http.StatusInternalServerError)
	}
	temp.Execute(w, nil)
}
