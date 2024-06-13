package main

import (
	"html/template"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

type team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var teams = []team{
	{ID: "1", Name: "Team Korea"},
	{ID: "2", Name: "Team USA"},
	{ID: "3", Name: "Team South Africa"},
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

func getTeams(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, teams)
}

func main() {
	router := gin.Default()
	router.GET("/teams", getTeams)
	router.Run("localhost:8000")
}
