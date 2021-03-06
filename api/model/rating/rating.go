package rating

import (
	"net/http"

	"../../model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection string = "rating"

type con struct {
	client   *mongo.Client
	database *mongo.Database
}

//Rating structure
type Rating struct {
	Tconst   string `bson:"tconst" json:"tconst"`
	Average  string `bson:"average" json:"average"`
	Numvotes string `bson:"numvotes" json:"numvotes"`
}

//RegisterHandler routes
func RegisterHandler(r *mux.Router, client *mongo.Client, database *mongo.Database) {
	condb := &con{client, database}

	//Route Handlers / Endpoints
	//r.HandleFunc("/api/movies/{tconst}", model.GetMovie).Methods("GET")
	r.HandleFunc("/api/"+collection, condb.Create).Methods("POST")
	r.HandleFunc("/api/"+collection+"/{tconst}", condb.Update).Methods("PUT")
	r.HandleFunc("/api/"+collection+"/{tconst}", condb.Delete).Methods("DELETE")
}

func getType() Rating {
	var rating Rating

	return rating
}

func getID() string {
	return "tconst"
}

//Create creates movie rating
func (mongocon *con) Create(w http.ResponseWriter, r *http.Request) {
	stype := getType()
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Create(db, w, r, collection, stype)
}

//Update updates movie rating
func (mongocon *con) Update(w http.ResponseWriter, r *http.Request) {
	stype := getType()
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Update(db, w, r, collection, stype, getID())
}

//Delete deletes movie rating
func (mongocon *con) Delete(w http.ResponseWriter, r *http.Request) {
	stype := getType()
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Delete(db, w, r, collection, stype, getID())
}
