package episode

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

var collection string = "episode"

func RegisterHandler(r *mux.Router, client *mongo.Client, database *mongo.Database) {
	condb := &con{client, database}

	//Route Handlers / Endpoints
	//r.HandleFunc("/api/movies/{tconst}", model.GetMovie).Methods("GET")
	r.HandleFunc("/api/"+collection, condb.CreateCrew).Methods("POST")
	r.HandleFunc("/api/"+collection+"/{tconst}", condb.UpdateCrew).Methods("PUT")
	r.HandleFunc("/api/"+collection+"/{tconst}", condb.DeleteCrew).Methods("DELETE")
}

//Crew structure
type Episode struct {
	tconst        string `bson:"tconst" json:"tconst"`
	parenttconst  string `bson:"parenttconst" json:"parenttconst"`
	seasonnumber  string `bson:"seasonnumber" json:"seasonnumber"`
	episodenumber string `bson:"episodenumber" json:"episodenumber"`
}

//CreateCrew creates movie crew
func (mongocon *con) CreateCrew(w http.ResponseWriter, r *http.Request) {
	var episode Episode
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Create(w, r, collection, episode, db)
}

//UpdateMovie updates movie data
func (mongocon *con) UpdateCrew(w http.ResponseWriter, r *http.Request) {
	var episode Episode
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Update(w, r, collection, episode, db)
}

//DeleteMovie deletes a movie
func (mongocon *con) DeleteCrew(w http.ResponseWriter, r *http.Request) {
	var episode Episode
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Delete(w, r, collection, episode, db)
}
