package crew

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

var collection string = "crew"

func RegisterHandler(r *mux.Router, client *mongo.Client, database *mongo.Database) {
	condb := &con{client, database}

	//Route Handlers / Endpoints
	//r.HandleFunc("/api/movies/{tconst}", model.GetMovie).Methods("GET")
	r.HandleFunc("/api/crew", condb.CreateCrew).Methods("POST")
	r.HandleFunc("/api/crew/{tconst}", condb.UpdateCrew).Methods("PUT")
	r.HandleFunc("/api/crew/{tconst}", condb.DeleteCrew).Methods("DELETE")
}

//Crew structure
type Crew struct {
	Tconst    string `bson:"tconst" json:"tconst"`       //alphanumeric unique identifier of the title
	Directors string `bson:"directors" json:"directors"` //(string array)
	Writers   string `bson:"writers" json:"writers"`     //(string array)
}

//CreateCrew creates movie crew
func (mongocon *con) CreateCrew(w http.ResponseWriter, r *http.Request) {
	var crew Crew
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Create(w, r, collection, crew, db)
}

//UpdateMovie updates movie data
func (mongocon *con) UpdateCrew(w http.ResponseWriter, r *http.Request) {
	var crew Crew
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Update(w, r, collection, crew, db)
}

//DeleteMovie deletes a movie
func (mongocon *con) DeleteCrew(w http.ResponseWriter, r *http.Request) {
	var crew Crew
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Delete(w, r, collection, crew, db)
}
