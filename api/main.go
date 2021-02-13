package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"model"
)

func main() {
	//init Router

	r := mux.NewRouter()

	//mock data
	model.Books = append(model.Books, Book{ID: "1", Title: "my title", Author: &Author{
		Firstname: "Tasos", Lastname: "Kar",
	}})
	model.Books = append(model.Books, Book{ID: "2", Title: "my title2", Author: &Author{
		Firstname: "Tasos2", Lastname: "Kar2",
	}})

	//Route Handlers / Endpoints
	r.HandleFunc("/api/books", model.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", model.GetBook).Methods("GET")
	r.HandleFunc("/api/books", model.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", model.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", model.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
