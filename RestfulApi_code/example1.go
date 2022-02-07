package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`	// json paths
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{
			Title: "Test title", 
			Desc:"Test Description", 
			Content:"Hello World"},
	}

	fmt.Println("Endpoint Hit: All Articlets Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func postAllArticlets(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Post Endpoint hit")
}


func runExample1(){

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", postAllArticlets).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", myRouter))
}
