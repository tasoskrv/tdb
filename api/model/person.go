package model

import (
	"net/http"
)

//Crew structure
type Person struct {
	nconst            string `bson:"nconst" json:"nconst"`
	primaryname       string `bson:"primaryname" json:"primaryname"`
	birthyear         string `bson:"birthyear" json:"birthyear"`
	deathyear         string `bson:"deathyear" json:"deathyear"`
	primaryproffesion string `bson:"primaryproffesion" json:"primaryproffesion"`
	knownfortitles    string `bson:"knownfortitles" json:"knownfortitles"`
}

//CreateCrew creates crew for movie
func (mongocon *MongoCon) CreatePerson(w http.ResponseWriter, r *http.Request) {

}

//UpdateCrew updates crew movie
func (mongocon *MongoCon) UpdatePerson(w http.ResponseWriter, r *http.Request) {

}

//DeleteCrew deletes movie crew
func (mongocon *MongoCon) DeletePerson(w http.ResponseWriter, r *http.Request) {

}
