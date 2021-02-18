package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"./model"
)

func Start() {
	r := mux.NewRouter()

	//mock data
	model.Books = append(model.Books, model.Book{ID: "1", Title: "my title", Author: &model.Author{
		Firstname: "Tasos", Lastname: "Kar",
	}})
	model.Books = append(model.Books, model.Book{ID: "2", Title: "my title2", Author: &model.Author{
		Firstname: "Tasos2", Lastname: "Kar2",
	}})

	//Route Handlers / Endpoints
	r.HandleFunc("/api/movies", model.GetMovies).Methods("GET")
	r.HandleFunc("/api/movies/{id}", model.GetMovie).Methods("GET")
	r.HandleFunc("/api/movies", model.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", model.UpdateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", model.DeleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
