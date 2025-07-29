package rote

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"myprojet/model"
	//"net/url"
)


func Artists(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found ", http.StatusNotFound)
		return
	}
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var artists []model.Artist

	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		log.Fatal(err)
	}
	if r.Method == http.MethodGet {
		input := "id"
		http.Redirect(w, r, "/artist/"+input, http.StatusSeeOther)
	}

	temp, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, "500 error internal server", http.StatusInternalServerError)
	}
	temp.Execute(w, artists)
}
