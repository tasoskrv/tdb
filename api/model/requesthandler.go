package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCon db related structure
type MongoCon struct {
	Client   *mongo.Client
	Database *mongo.Database
}

//Get inserts document
func Get(dbcon *MongoCon, w http.ResponseWriter, r *http.Request, c string, id string) {
	SetHeaders(w)
	params := mux.Vars(r) //Get params

	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := dbcon.Database.Collection(c)

	var results []bson.M
	cur, err := collection.Find(
		context.Background(),
		bson.M{"tconst": params[id]},
		options.Find().SetProjection(bson.M{"_id": 0}),
	)

	defer cur.Close(context.Background())

	if err != nil {
		RespondErr(w, r, http.StatusConflict, err)
		return
	}

	cur.All(context.Background(), &results)
	Respond(w, r, http.StatusOK, results[0])
}

//Create inserts document
func Create(dbcon *MongoCon, w http.ResponseWriter, r *http.Request, c string, v interface{}) {
	SetHeaders(w)

	_ = DecodeBody(r, &v)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := dbcon.Database.Collection(c)
	result, err := collection.InsertOne(ctx, v)

	if err != nil {
		RespondErr(w, r, http.StatusConflict, err)
		return
	}
	_ = result.InsertedID.(primitive.ObjectID).Hex()

	Respond(w, r, http.StatusOK, v)
}

//Update updates document
func Update(dbcon *MongoCon, w http.ResponseWriter, r *http.Request, c string, v interface{}, id string) {
	SetHeaders(w)
	params := mux.Vars(r) //Get params

	_ = DecodeBody(r, &v)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := dbcon.Database.Collection(c)

	_, err := collection.UpdateOne(
		ctx,
		bson.M{id: params[id]},
		bson.M{"$set": v},
	)

	if err != nil {
		RespondErr(w, r, http.StatusConflict, err)
		return
	}

	Respond(w, r, http.StatusOK, v)
}

//Delete removes document
func Delete(dbcon *MongoCon, w http.ResponseWriter, r *http.Request, c string, v interface{}, id string) {
	SetHeaders(w)
	params := mux.Vars(r) //Get params

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := dbcon.Database.Collection(c)

	_, err := collection.DeleteOne(ctx, bson.M{id: params[id]})
	if err != nil {
		RespondErr(w, r, http.StatusConflict, err)
	}

	Respond(w, r, http.StatusOK, nil)
}

//SetHeaders adds headers to http request
func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

//DecodeBody returns decoded data from http request (raw data)
func DecodeBody(r *http.Request, v interface{}) error {
	if r.Body != nil {
		defer r.Body.Close()
		return json.NewDecoder(r.Body).Decode(v)
	}
	return errors.New("Undefined body")
}

//EncodeBody returns encoded data
func EncodeBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

//Respond builds successful response
func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	SetHeaders(w)
	w.WriteHeader(status)
	if data != nil {
		EncodeBody(w, r, data)
	}
}

//RespondErr builds error response
func RespondErr(w http.ResponseWriter, r *http.Request, status int, args ...interface{}) {
	Respond(w, r, status, map[string]interface{}{
		"error": map[string]interface{}{
			"message": fmt.Sprint(args...),
		},
	})
}

//RespondHTTPErr builds http error response
func RespondHTTPErr(w http.ResponseWriter, r *http.Request, status int) {
	RespondErr(w, r, status, http.StatusText(status))
}
