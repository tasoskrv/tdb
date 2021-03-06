package api

import (
	"log"
	"net/http"

	"./model/crew"
	"./model/movie"
	"./model/person"
	"./model/rating"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

//Start main.DBconn ...
func Start(client *mongo.Client, database *mongo.Database) {
	r := mux.NewRouter()
	movie.RegisterHandler(r, client, database)
	crew.RegisterHandler(r, client, database)
	rating.RegisterHandler(r, client, database)
	person.RegisterHandler(r, client, database)

	log.Fatal(http.ListenAndServe(":8000", r))
}
