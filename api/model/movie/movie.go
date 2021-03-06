package movie

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

var collection string = "movie"

func RegisterHandler(r *mux.Router, client *mongo.Client, database *mongo.Database) {
	condb := &con{client, database}

	//Route Handlers / Endpoints
	//r.HandleFunc("/api/movies/{tconst}", model.GetMovie).Methods("GET")
	r.HandleFunc("/api/movie", condb.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{tconst}", condb.UpdateMovie).Methods("PUT")
	r.HandleFunc("/api/movie/{tconst}", condb.DeleteMovie).Methods("DELETE")
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

//CreateMovie creates a movie
func (mongocon *con) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Create(w, r, collection, movie, db)
}

//UpdateMovie updates movie data
func (mongocon *con) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Update(w, r, collection, movie, db)
}

//DeleteMovie deletes a movie
func (mongocon *con) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	db := &model.MongoCon{Database: mongocon.database, Client: mongocon.client}
	model.Delete(w, r, collection, movie, db)
}
