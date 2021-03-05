package model

import (
	"net/http"
)

//Crew structure
type Episode struct {
	tconst        string `bson:"tconst" json:"tconst"`
	parenttconst  string `bson:"parenttconst" json:"parenttconst"`
	seasonnumber  string `bson:"seasonnumber" json:"seasonnumber"`
	episodenumber string `bson:"episodenumber" json:"episodenumber"`
}

//CreateCrew creates crew for movie
func (mongocon *MongoCon) CreateEpisode(w http.ResponseWriter, r *http.Request) {

}

//UpdateCrew updates crew movie
func (mongocon *MongoCon) UpdateEpisode(w http.ResponseWriter, r *http.Request) {

}

//DeleteCrew deletes movie crew
func (mongocon *MongoCon) DeleteEpisode(w http.ResponseWriter, r *http.Request) {

}
