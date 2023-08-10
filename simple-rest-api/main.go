package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

var articles []Article

func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", getArticles)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "API Homepage")
}

func main() {
	addData()
	handleRequests()
}

func addData() {
	articles = append(articles, Article{
		Title:       "Now or Never",
		Author:      "Vishal Singh",
		Description: "Just try and cry",
	})
}
