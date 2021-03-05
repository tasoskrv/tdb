package model

import (
	"net/http"
)

//Crew structure
type Crew struct {
	Tconst    string `bson:"tconst" json:"tconst"`       //alphanumeric unique identifier of the title
	Directors string `bson:"directors" json:"directors"` //(string array)
	Writers   string `bson:"writers" json:"writers"`     //(string array)
}

//CreateCrew creates crew for movie
func (mongocon *MongoCon) CreateCrew(w http.ResponseWriter, r *http.Request) {

}

//UpdateCrew updates crew movie
func (mongocon *MongoCon) UpdateCrew(w http.ResponseWriter, r *http.Request) {

}

//DeleteCrew deletes movie crew
func (mongocon *MongoCon) DeleteCrew(w http.ResponseWriter, r *http.Request) {

}
