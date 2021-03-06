package movie

import (
	"net/http"

	"../../model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection string = "movie"

type con struct {
	client   *mongo.Client
	database *mongo.Database
}

//Movie structure
type Movie struct {
	Tconst         string `bson:"tconst" json:"tconst"`                 //alphanumeric unique identifier of the title
	TitleType      string `bson:"titletype" json:"titletype"`           //the type/format of the title (e.g. movie, short, tvseries, tvepisode, video, etc)
	PrimaryTitle   string `bson:"primarytitle" json:"primarytitle"`     //the more popular title / the title used by the filmmakers on promotional materials at the point of release
	OriginalTitle  string `bson:"originaltitle" json:"originaltitle"`   //original title, in the original language
	IsAdult        bool   `bson:"isadult" json:"isadult"`               //0: non-adult title; 1: adult title
	StartYear      int    `bson:"startyear" json:"startyear"`           //represents the release year of a title. In the case of TV Series, it is the series start year
	EndYear        int    `bson:"endyear" json:"endyear"`               //TV Series end year. ‘\N’ for all other title types
	RuntimeMinutes int    `bson:"runtimeminutes" json:"runtimeminutes"` //primary runtime of the title, in minutes
	Genres         string `bson:"genres" json:"genres"`
}

//RegisterHandler routes Handlers / Endpoints
func RegisterHandler(r *mux.Router, client *mongo.Client, database *mongo.Database) {
	condb := &con{client, database}

	//r.HandleFunc("/api/movies/{tconst}", model.GetMovie).Methods("GET")
	r.HandleFunc("/api/"+collection, condb.Create).Methods("POST")
	r.HandleFunc("/api/"+collection+"/{tconst}", condb.Update).Methods("PUT")
	r.HandleFunc("/api/"+collection+"/{tconst}", condb.Delete).Methods("DELETE")
}

func getType() Movie {
	var movie Movie

	return movie
}

func getID() string {
	return "tconst"
}

//Create creates a movie
func (mongocon *con) Create(w http.ResponseWriter, r *http.Request) {
	stype := getType()
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Create(db, w, r, collection, stype)
}

//Update updates movie data
func (mongocon *con) Update(w http.ResponseWriter, r *http.Request) {
	stype := getType()
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Update(db, w, r, collection, stype, getID())
}

//Delete deletes a movie
func (mongocon *con) Delete(w http.ResponseWriter, r *http.Request) {
	stype := getType()
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Delete(db, w, r, collection, stype, getID())
}
