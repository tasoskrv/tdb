package model

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	AverageRating  float32 `bson:"averagerating" json:"averagerating"`   //(float)
	NumVotes       int     `bson:"numvotes" json:"numvotes"`             //(int)
	Directors      string  `bson:"directors" json:"directors"`           //(string array)
	Actors         string  `bson:"actors" json:"actors"`                 //(string array)
}

//MongoCon structure
type MongoCon struct {
	Client   *mongo.Client
	Database *mongo.Database
}

type response struct {
	Success bool
	ID      string
	Data    Movie
}

//CreateMovie ... creates a movie
func (mongocon *MongoCon) CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&movie)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	mCol := mongocon.Database.Collection("movie")
	result, err2 := mCol.InsertOne(ctx, bson.M{
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

	if err2 != nil {
		fmt.Println(err2)
	}

	res := response{
		Success: true,
		ID:      mid,
		Data:    movie,
	}

	b, err := json.Marshal(res)
	if err != nil {
		// Handle Error
	}
	w.Write(b)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	/*
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r) //Get params
		for index, item := range Books {
			if item.ID == params["id"] {
				Books = append(Books[:index], Books[index+1:]...)

				var book Book
				_ = json.NewDecoder(r.Body).Decode(&book)
				book.ID = strconv.Itoa(rand.Intn(1000000))
				Books = append(Books, book)
				json.NewEncoder(w).Encode(book)
				return
			}
		}

		json.NewEncoder(w).Encode(Books)
	*/
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	/*
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r) //Get params
		for index, item := range Books {
			if item.ID == params["id"] {
				Books = append(Books[:index], Books[index+1:]...)
				break
			}
		}

		json.NewEncoder(w).Encode(Books)*/
}
