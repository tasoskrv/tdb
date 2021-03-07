package crew

import (
	"net/http"

	"../../model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection string = "crew"

type con struct {
	client   *mongo.Client
	database *mongo.Database
}

//Crew structure
type Crew struct {
	Tconst    string `bson:"tconst" json:"tconst"`       //alphanumeric unique identifier of the title
	Directors string `bson:"directors" json:"directors"` //(string array)
	Writers   string `bson:"writers" json:"writers"`     //(string array)
}

//RegisterHandler routes Handlers / Endpoints
func RegisterHandler(r *mux.Router, client *mongo.Client, database *mongo.Database) {
	condb := &con{client, database}

	r.HandleFunc("/api/"+collection+"/{tconst}", condb.Get).Methods("GET")
	r.HandleFunc("/api/"+collection, condb.Create).Methods("POST")
	r.HandleFunc("/api/"+collection+"/{tconst}", condb.Update).Methods("PUT")
	r.HandleFunc("/api/"+collection+"/{tconst}", condb.Delete).Methods("DELETE")
}

func getType() Crew {
	var crew Crew

	return crew
}

func getID() string {
	return "tconst"
}

//Get return single movie
func (mongocon *con) Get(w http.ResponseWriter, r *http.Request) {
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Get(db, w, r, collection, getID())
}

//Create creates movie crew
func (mongocon *con) Create(w http.ResponseWriter, r *http.Request) {
	stype := getType()
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Create(db, w, r, collection, stype)
}

//Update updates movie crew
func (mongocon *con) Update(w http.ResponseWriter, r *http.Request) {
	stype := getType()
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Update(db, w, r, collection, stype, getID())
}

//Delete deletes movie crew
func (mongocon *con) Delete(w http.ResponseWriter, r *http.Request) {
	stype := getType()
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Delete(db, w, r, collection, stype, getID())
}
