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

	modelStruct := &model.MongoCon{Client: client, Database: database}

	r := mux.NewRouter()
	//Route Handlers / Endpoints
	//r.HandleFunc("/api/movies", model.GetMovies).Methods("GET")
	//r.HandleFunc("/api/movies/{id}", model.GetMovie).Methods("GET")
	r.HandleFunc("/api/movies", modelStruct.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies/{tconst}", modelStruct.UpdateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{tconst}", modelStruct.DeleteMovie).Methods("DELETE")

	r.HandleFunc("/api/crew", modelStruct.CreateCrew).Methods("POST")
	r.HandleFunc("/api/crew/{tconst}", modelStruct.UpdateCrew).Methods("PUT")
	r.HandleFunc("/api/crew/{tconst}", modelStruct.DeleteCrew).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
