package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title 	string `json:"Title"`
	Desc 	string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome to the HomePage!")
	fmt.Println("Endpoint hit: homePage")
}

func handleRequests() {
	// routes
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)

	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
	}

	handleRequests()
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllArticles")

	// encodes Articles array into a JSON string
	// writes as part of response
	json.NewEncoder(w).Encode(Articles)
}
