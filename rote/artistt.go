package rote

import (
	"html/template"
	"net/http"
	"strconv"

	"myprojet/artists"
)

func Artist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	ListeAtists, err := artists.GetArtists()
	if err != nil {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/artist/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	for _, artist := range ListeAtists {
		if id == artist.ID {
			tmpl, err := template.ParseFiles("template/result.html")
			if err != nil {
				http.Error(w, "Error loading page", http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, artist)
			return
		}
	}

	http.Error(w, "404 artist not found", http.StatusNotFound)
}
