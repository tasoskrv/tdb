package api

import (
	"log"
	"net/http"

	"./model"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gorilla/mux"
)

//Start main.DBconn ...
func Start(cl *mongo.Client, d *mongo.Database) {

	model := &model.Users{Cl: cl, DB: d}

	r := mux.NewRouter()
	//Route Handlers / Endpoints
	//r.HandleFunc("/api/movies", model.GetMovies).Methods("GET")
	//r.HandleFunc("/api/movies/{id}", model.GetMovie).Methods("GET")
	r.HandleFunc("/api/movies", model.CreateMovie).Methods("POST")
	//r.HandleFunc("/api/movies/{id}", model.UpdateMovie).Methods("PUT")
	//r.HandleFunc("/api/movies/{id}", model.DeleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
