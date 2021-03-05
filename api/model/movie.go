package model

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

//CreateMovie ... creates a movie
func (mongocon *MongoCon) CreateMovie(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	var movie Movie

	_ = DecodeBody(r, &movie)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	mCol := mongocon.Database.Collection("movie")
	result, err := mCol.InsertOne(ctx, bson.M{
		"tconst":         movie.Tconst,
		"titletype":      movie.TitleType,
		"primarytitle":   movie.PrimaryTitle,
		"originaltitle":  movie.OriginalTitle,
		"isadult":        movie.IsAdult,
		"endyear":        movie.EndYear,
		"runtimeminutes": movie.RuntimeMinutes,
		"genres":         movie.Genres,
	},
	)

	_ = result.InsertedID.(primitive.ObjectID).Hex()

	if err != nil {
		RespondErr(w, r, http.StatusConflict, err)
		return
	}

	Respond(w, r, http.StatusOK, movie)
}

//UpdateMovie updates movie data
func (mongocon *MongoCon) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	params := mux.Vars(r) //Get params

	var movie Movie
	_ = DecodeBody(r, &movie)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	moviesCollection := mongocon.Database.Collection("movie")

	_, err := moviesCollection.UpdateOne(
		ctx,
		bson.M{"tconst": params["tconst"]},
		bson.M{"$set": movie},
	)

	if err != nil {
		RespondErr(w, r, http.StatusConflict, err)
		return
	}

	Respond(w, r, http.StatusOK, movie)
}

//DeleteMovie deletes a movie
func (mongocon *MongoCon) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	params := mux.Vars(r) //Get params

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	moviesCollection := mongocon.Database.Collection("movie")

	_, err := moviesCollection.DeleteOne(ctx, bson.M{"tconst": params["tconst"]})
	if err != nil {
		RespondErr(w, r, http.StatusConflict, err)
	}

	Respond(w, r, http.StatusOK, nil)
}
