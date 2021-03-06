package rating

import (
	"net/http"

	"../../model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type con struct {
	client   *mongo.Client
	database *mongo.Database
}

var collection string = "rating"

func RegisterHandler(r *mux.Router, client *mongo.Client, database *mongo.Database) {
	condb := &con{client, database}

	//Route Handlers / Endpoints
	//r.HandleFunc("/api/movies/{tconst}", model.GetMovie).Methods("GET")
	r.HandleFunc("/api/rating", condb.CreateCrew).Methods("POST")
	r.HandleFunc("/api/rating/{tconst}", condb.UpdateCrew).Methods("PUT")
	r.HandleFunc("/api/rating/{tconst}", condb.DeleteCrew).Methods("DELETE")
}

//Crew structure
type Rating struct {
	Tconst   string `bson:"tconst" json:"tconst"`
	average  string `bson:"average" json:"average"`
	numvotes string `bson:"numvotes" json:"numvotes"`
}

//CreateCrew creates movie crew
func (mongocon *con) CreateCrew(w http.ResponseWriter, r *http.Request) {
	var rating Rating
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Create(w, r, collection, rating, db)
}

//UpdateMovie updates movie data
func (mongocon *con) UpdateCrew(w http.ResponseWriter, r *http.Request) {
	var rating Rating
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Update(w, r, collection, rating, db)
}

//DeleteMovie deletes a movie
func (mongocon *con) DeleteCrew(w http.ResponseWriter, r *http.Request) {
	var rating Rating
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Delete(w, r, collection, rating, db)
}
