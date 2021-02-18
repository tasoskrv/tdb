package model

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Movie struct {
	Tconst         string `bson:"tconst" json:"tconst"`                 //alphanumeric unique identifier of the title
	TitleType      string `bson:"titletype" json:"titletype"`           //the type/format of the title (e.g. movie, short, tvseries, tvepisode, video, etc)
	PrimaryTitle   string `bson:"primarytitle" json:"primarytitle"`     //the more popular title / the title used by the filmmakers on promotional materials at the point of release
	OriginalTitle  string `bson:"originaltitle" json:"originaltitle"`   //original title, in the original language
	IsAdult        bool   `bson:"isadult" json:"isadult"`               //0: non-adult title; 1: adult title
	StartYear      int    `bson:"startyear" json:"startyear"`           //represents the release year of a title. In the case of TV Series, it is the series start year
	EndYear        int    `bson:"endyear" json:"endyear"`               //TV Series end year. ‘\N’ for all other title types
	RuntimeMinutes int    `bson:"runtimeminutes" json:"runtimeminutes"` //primary runtime of the title, in minutes
	Genres         string `bson:"genres" json:"genres"`                 //(string array) – includes up to three genres associated with the title

}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params

	for _, item := range Books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&movie)

	dbcon, f, _ := InitDatabase()

	mCol := dbcon.db.Collection("movie")
	_, err2 := mCol.InsertOne(dbcon.ctx, bson.D{
		{Key: "tconst", Value: movie.Tconst},
		{Key: "titletype", Value: movie.TitleType},
	})
	/*
		mCol.InsertOne(dbcon.ctx, Movie{
			Tconst:         movie.Tconst,
			TitleType:      movie.TitleType,
			PrimaryTitle:   movie.PrimaryTitle,
			OriginalTitle:  movie.OriginalTitle,
			IsAdult:        movie.IsAdult,
			EndYear:        movie.EndYear,
			RuntimeMinutes: movie.RuntimeMinutes,
			Genres:         movie.Genres,
		})
	*/
	if err2 != nil {
		fmt.Println(err2)
	}

	f()

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
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
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	for index, item := range Books {
		if item.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(Books)
}

type dbconn struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
}

func InitDatabase() (dbconn, func(), error) {
	connMongo := "mongodb://127.0.0.1:2717"

	dbMongo := "tdb"

	//Create client
	client, err := mongo.NewClient(options.Client().ApplyURI(connMongo))
	if err != nil {
		return dbconn{}, nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	//Connect to mongo
	err = client.Connect(ctx)
	if err != nil {
		return dbconn{}, nil, err
	}

	// Return disconnect as func to defer in main
	f := func() {
		defer cancel()
		client.Disconnect(ctx)
	}

	//Ping Database
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return dbconn{}, f, err
	}
	db := client.Database(dbMongo)

	return dbconn{client: client, db: db, ctx: ctx}, f, nil

}

/*
func connectMysql() {
	//  does not open any physical connection to the database server, but it will validate its arguments
	db, err := sql.Open("mysql", "root:mysql123@tcp(127.0.0.1:3306)/tdb")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	} else {
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}
	}


	//res, err := db.Exec(`insert into movie (tconst, titletype, primarytitle, originaltitle, isadult, startyear, endyear, runtimeminutes, genres)
	//values ('` + tconst + `')`)

}
*/
