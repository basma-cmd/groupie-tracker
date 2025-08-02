package rote

import (
	"html/template"
	"myprojet/artists"
	"net/http"
	
)


func Artists(w http.ResponseWriter, r *http.Request) {


	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found ", http.StatusNotFound)
		return
	}
	artists, err := artists.GetArtists()
	if err != nil {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	temp, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "500 error internal server", http.StatusInternalServerError)
	}
	temp.Execute(w, artists)
	
}
