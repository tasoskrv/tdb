package model

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	
	ID int	`json:"id"`
	Tconst string `json:"tconst"`//alphanumeric unique identifier of the title
	TitleType `json:"titletype"`//
	`json:""`
	`json:""`
	`json:""`
	`json:""`
	`json:""`
	`json:""`
	`json:""`
	`json:""`

	titleType (string) – the type/format of the title (e.g. movie, short, tvseries, tvepisode, video, etc)
	primaryTitle (string) – the more popular title / the title used by the filmmakers on promotional materials at the point of release
	originalTitle (string) - original title, in the original language
	isAdult (boolean) - 0: non-adult title; 1: adult title
	startYear (YYYY) – represents the release year of a title. In the case of TV Series, it is the series start year
	endYear (YYYY) – TV Series end year. ‘\N’ for all other title types
	runtimeMinutes – primary runtime of the title, in minutes
	genres (string array) – includes up to three genres associated with the title	
}


//Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author ...
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Books ...
var Books []Book

//GetBooks ...
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
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

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	Books = append(Books, book)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
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

func DeleteBook(w http.ResponseWriter, r *http.Request) {
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
