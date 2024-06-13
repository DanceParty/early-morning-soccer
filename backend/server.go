package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type Team struct {
	Name   string
	Wins   int
	Losses int
	Draws  int
}

func showTeam(w http.ResponseWriter, r *http.Request) {
	team := Team{"Team Korea", 1, 4, 0}
	fp := path.Join("views", "index.html")
	temp, err := template.ParseFiles(fp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := temp.Execute(w, team); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", showTeam)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
