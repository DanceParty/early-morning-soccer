package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func showTeam(w http.ResponseWriter, r *http.Request) {
	team := team{"1", "Team Korea"}
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
