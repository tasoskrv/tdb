package model

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Movie structure
type Movie struct {
	Tconst         string  `bson:"tconst" json:"tconst"`                 //alphanumeric unique identifier of the title
	TitleType      string  `bson:"titletype" json:"titletype"`           //the type/format of the title (e.g. movie, short, tvseries, tvepisode, video, etc)
	PrimaryTitle   string  `bson:"primarytitle" json:"primarytitle"`     //the more popular title / the title used by the filmmakers on promotional materials at the point of release
	OriginalTitle  string  `bson:"originaltitle" json:"originaltitle"`   //original title, in the original language
	IsAdult        bool    `bson:"isadult" json:"isadult"`               //0: non-adult title; 1: adult title
	StartYear      int     `bson:"startyear" json:"startyear"`           //represents the release year of a title. In the case of TV Series, it is the series start year
	EndYear        int     `bson:"endyear" json:"endyear"`               //TV Series end year. ‘\N’ for all other title types
	RuntimeMinutes int     `bson:"runtimeminutes" json:"runtimeminutes"` //primary runtime of the title, in minutes
	Genres         string  `bson:"genres" json:"genres"`
	AverageRating  float32 `bson:"averagerating" json:"averagerating"` //(float)
	NumVotes       int     `bson:"numvotes" json:"numvotes"`           //(int)
	Directors      string  `bson:"directors" json:"directors"`         //(string array)
	Actors         string  `bson:"actors" json:"actors"`               //(string array)
}

//MongoCon structure
type MongoCon struct {
	Client   *mongo.Client
	Database *mongo.Database
}

type response struct {
	success bool
	id      string
	data    Movie
}

type ResponseDelete struct {
	Success bool
}

//CreateMovie ... creates a movie
func (mongocon *MongoCon) CreateMovie(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	var movie Movie

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&movie)

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

	mid := result.InsertedID.(primitive.ObjectID).Hex()

	if err != nil {
		writeErrorLogs(err)
	}

	res := response{
		success: true,
		id:      mid,
		data:    movie,
	}

	b, errR := json.Marshal(res)
	writeResponse(w, b, errR)
}

//UpdateMovie updates movie data
func (mongocon *MongoCon) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	params := mux.Vars(r) //Get params

	var movie Movie

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&movie)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	moviesCollection := mongocon.Database.Collection("movie")

	_, err := moviesCollection.UpdateOne(
		ctx,
		bson.M{"tconst": params["tconst"]},
		bson.D{
			{"$set", bson.D{{"titletype", movie.TitleType}}},
			{"$set", bson.D{{"primarytitle", movie.PrimaryTitle}}},
			{"$set", bson.D{{"originaltitle", movie.OriginalTitle}}},
			{"$set", bson.D{{"isadult", movie.IsAdult}}},
			{"$set", bson.D{{"startyear", movie.StartYear}}},
			{"$set", bson.D{{"endyear", movie.EndYear}}},
			{"$set", bson.D{{"runtimeminutes", movie.RuntimeMinutes}}},
			{"$set", bson.D{{"genres", movie.Genres}}},
		},
	)
	if err != nil {
		writeErrorLogs(err)
	}
	//fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

	res := response{
		success: true,
		data:    movie,
	}

	b, errR := json.Marshal(res)

	writeResponse(w, b, errR)
}

type bot interface {
	DeleteMovie(w http.ResponseWriter, r *http.Request)
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
		writeErrorLogs(err)
	}

	res := ResponseDelete{
		Success: true,
	}

	b, _ := json.Marshal(res)

	w.Write(b)
	//writeResponse(w, b, errR)
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func writeResponse(w http.ResponseWriter, b []byte, err error) {
	if err != nil {
		// Handle Error
		w.Write([]byte(err.Error()))
	}

	w.Write(b)
}

func writeErrorLogs(err error) {
	log.Fatal(err)
}
