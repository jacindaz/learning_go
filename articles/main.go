package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
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
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
	}

	handleRequests()
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")

	// encodes Articles array into a JSON string
	// writes as part of response
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleArticle")

	vars := mux.Vars(r)
	key := vars["id"]

	index, err := strconv.Atoi(key)
	if err == nil {
		article := Articles[index-1]
		json.NewEncoder(w).Encode(article)
	} else {
		fmt.Fprintf(w, "Error, could not process request")
	}
}
