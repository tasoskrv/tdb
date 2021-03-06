package person

import (
	"net/http"

	"../../model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection string = "person"

type con struct {
	client   *mongo.Client
	database *mongo.Database
}

//Person structure
type Person struct {
	Nconst            string `bson:"nconst" json:"nconst"`
	Primaryname       string `bson:"primaryname" json:"primaryname"`
	Birthyear         string `bson:"birthyear" json:"birthyear"`
	Deathyear         string `bson:"deathyear" json:"deathyear"`
	Primaryproffesion string `bson:"primaryproffesion" json:"primaryproffesion"`
	Knownfortitles    string `bson:"knownfortitles" json:"knownfortitles"`
}

//RegisterHandler routes Handlers / Endpoints
func RegisterHandler(r *mux.Router, client *mongo.Client, database *mongo.Database) {
	condb := &con{client, database}

	//r.HandleFunc("/api/movies/{tconst}", model.GetMovie).Methods("GET")
	r.HandleFunc("/api/"+collection, condb.Create).Methods("POST")
	r.HandleFunc("/api/"+collection+"/{tconst}", condb.Update).Methods("PUT")
	r.HandleFunc("/api/"+collection+"/{tconst}", condb.Delete).Methods("DELETE")
}

func getType() Person {
	var person Person

	return person
}

func getID() string {
	return "nconst"
}

//Create creates person
func (mongocon *con) Create(w http.ResponseWriter, r *http.Request) {
	stype := getType()
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Create(db, w, r, collection, stype)
}

//Update updates person
func (mongocon *con) Update(w http.ResponseWriter, r *http.Request) {
	stype := getType()
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Update(db, w, r, collection, stype, getID())
}

//Delete deletes person
func (mongocon *con) Delete(w http.ResponseWriter, r *http.Request) {
	stype := getType()
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Delete(db, w, r, collection, stype, getID())
}
