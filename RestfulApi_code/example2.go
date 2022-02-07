package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct{
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`		// when json is a struct, use *
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

type Books []Book;

// Mock Data
var books Books;


// handlers
// get all books
func getBooks(w http.ResponseWriter, r *http.Request){		// similar to app.get("", function(req, res))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)	
	fmt.Println("getting books...");
}
// get single book with id
func getBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r) 									// get any parameter from request
	fmt.Println("getting the book with id " + params["id"]);

	w.Header().Set("Content-Type", "application/json")
	// loop through books and find an id
	for _,item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return;
		}
	}
	json.NewEncoder(w).Encode(Book{})	
	return;
}
// create book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	_ = json.NewDecoder(r.Body).Decode(&newBook);
	fmt.Printf("Got a new book: %v", newBook)
	newBook.ID = strconv.Itoa(rand.Intn(10000000));		// note: this is not safe
	books = append(books, newBook)
	json.NewEncoder(w).Encode(newBook)	
	fmt.Println();
	fmt.Printf("New Book Collection: %v", books)
	fmt.Println();
}
// update a book
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var updatedBook Book;
	_ = json.NewDecoder(r.Body).Decode(&updatedBook)
	for index, item := range books {
		if item.ID == params["id"]{
			books[index] = updatedBook
			books[index].ID = params["id"]
			fmt.Printf("Updating the book with id " + params["id"])
			fmt.Println()
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

// delete book
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"]{
			fmt.Printf("Deleting the book with id "+params["id"])
			fmt.Println()
			books = append(books[:index], books[index+1:]... )		// remove the indexed book
			break
		}
	}
	fmt.Printf("Sending back current book collection...")
	fmt.Println()
	json.NewEncoder(w).Encode(books)
}

// initializer
func runExample2() {
	books = append(books, 
		Book{
			ID:"1", 
			ISBN:"1111", 
			Title:"title1", 
			Author: &Author{Firstname:"first1", Lastname:"last1"}},
		Book{
			ID:"2", 
			ISBN:"2222", 
			Title:"title2", 
			Author: &Author{Firstname:"first2", Lastname:"last2"}},
		Book{
			ID:"3", 
			ISBN:"3333", 
			Title:"title3", 
			Author: &Author{Firstname:"first3", Lastname:"last3"}},
		Book{
			ID:"4", 
			ISBN:"4444", 
			Title:"title4", 
			Author: &Author{Firstname:"first4", Lastname:"last4"}},
	)

	// init router
	r := mux.NewRouter().StrictSlash(true);

	// end points
	r.HandleFunc("/", homePage)
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// start the server
	log.Fatal(http.ListenAndServe(":8001", r))
	// note: you can only use one localhost in a go project I guess?
	// :8000 doesn't work
}
