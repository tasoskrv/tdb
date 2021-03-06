package episode

import (
	"net/http"

	"../../model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection string = "episode"

type con struct {
	client   *mongo.Client
	database *mongo.Database
}

//Episode structure
type Episode struct {
	Tconst        string `bson:"tconst" json:"tconst"`
	Parenttconst  string `bson:"parenttconst" json:"parenttconst"`
	Seasonnumber  string `bson:"seasonnumber" json:"seasonnumber"`
	Episodenumber string `bson:"episodenumber" json:"episodenumber"`
}

//RegisterHandler routes Handlers / Endpoints
func RegisterHandler(r *mux.Router, client *mongo.Client, database *mongo.Database) {
	condb := &con{client, database}

	//r.HandleFunc("/api/movies/{tconst}", model.GetMovie).Methods("GET")
	r.HandleFunc("/api/"+collection, condb.Create).Methods("POST")
	r.HandleFunc("/api/"+collection+"/{tconst}", condb.Update).Methods("PUT")
	r.HandleFunc("/api/"+collection+"/{tconst}", condb.Delete).Methods("DELETE")
}

func getType() Episode {
	var episode Episode

	return episode
}

func getID() string {
	return "tconst"
}

//Create creates movie episode
func (mongocon *con) Create(w http.ResponseWriter, r *http.Request) {
	stype := getType()
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Create(db, w, r, collection, stype)
}

//Update updates movie episode
func (mongocon *con) Update(w http.ResponseWriter, r *http.Request) {
	stype := getType()
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Update(db, w, r, collection, stype, getID())
}

//Delete deletes movie episode
func (mongocon *con) Delete(w http.ResponseWriter, r *http.Request) {
	stype := getType()
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Delete(db, w, r, collection, stype, getID())
}
