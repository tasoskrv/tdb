package model

import (
	"net/http"
)

//Crew structure
type Rating struct {
	Tconst   string `bson:"tconst" json:"tconst"`
	average  string `bson:"average" json:"average"`
	numvotes string `bson:"numvotes" json:"numvotes"`
}

//CreateCrew creates crew for movie
func (mongocon *MongoCon) CreateRating(w http.ResponseWriter, r *http.Request) {

}

//UpdateCrew updates crew movie
func (mongocon *MongoCon) UpdateRating(w http.ResponseWriter, r *http.Request) {

}

//DeleteCrew deletes movie crew
func (mongocon *MongoCon) DeleteRating(w http.ResponseWriter, r *http.Request) {

}
