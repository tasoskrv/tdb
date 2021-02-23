package api

import (
	"log"
	"net/http"

	"./model"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gorilla/mux"
)

//Start main.DBconn ...
func Start(client *mongo.Client, database *mongo.Database) {

	modelStr := &model.MongoCon{Client: client, Database: database}

	r := mux.NewRouter()
	//Route Handlers / Endpoints
	//r.HandleFunc("/api/movies", model.GetMovies).Methods("GET")
	//r.HandleFunc("/api/movies/{id}", model.GetMovie).Methods("GET")
	r.HandleFunc("/api/movies", modelStr.CreateMovie).Methods("POST")
	//r.HandleFunc("/api/movies/{id}", model.UpdateMovie).Methods("PUT")
	//r.HandleFunc("/api/movies/{id}", model.DeleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
