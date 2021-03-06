package person

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

var collection string = "person"

func RegisterHandler(r *mux.Router, client *mongo.Client, database *mongo.Database) {
	condb := &con{client, database}

	//Route Handlers / Endpoints
	//r.HandleFunc("/api/movies/{tconst}", model.GetMovie).Methods("GET")
	r.HandleFunc("/api/"+collection, condb.CreateCrew).Methods("POST")
	r.HandleFunc("/api/"+collection+"/{tconst}", condb.UpdateCrew).Methods("PUT")
	r.HandleFunc("/api/"+collection+"/{tconst}", condb.DeleteCrew).Methods("DELETE")
}

//Crew structure
type Person struct {
	nconst            string `bson:"nconst" json:"nconst"`
	primaryname       string `bson:"primaryname" json:"primaryname"`
	birthyear         string `bson:"birthyear" json:"birthyear"`
	deathyear         string `bson:"deathyear" json:"deathyear"`
	primaryproffesion string `bson:"primaryproffesion" json:"primaryproffesion"`
	knownfortitles    string `bson:"knownfortitles" json:"knownfortitles"`
}

//CreateCrew creates movie crew
func (mongocon *con) CreateCrew(w http.ResponseWriter, r *http.Request) {
	var person Person
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Create(w, r, collection, person, db)
}

//UpdateMovie updates movie data
func (mongocon *con) UpdateCrew(w http.ResponseWriter, r *http.Request) {
	var person Person
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Update(w, r, collection, person, db)
}

//DeleteMovie deletes a movie
func (mongocon *con) DeleteCrew(w http.ResponseWriter, r *http.Request) {
	var person Person
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Delete(w, r, collection, person, db)
}
