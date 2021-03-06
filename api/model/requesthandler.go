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
)

//Create func
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

/*****************************************/

func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func DecodeBody(r *http.Request, v interface{}) error {
	if r.Body != nil {
		defer r.Body.Close()
		return json.NewDecoder(r.Body).Decode(v)
	}
	return errors.New("Undefined body")
}

func EncodeBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if data != nil {
		EncodeBody(w, r, data)
	}
}

func RespondErr(w http.ResponseWriter, r *http.Request, status int, args ...interface{}) {
	Respond(w, r, status, map[string]interface{}{
		"error": map[string]interface{}{
			"message": fmt.Sprint(args...),
		},
	})
}

func RespondHTTPErr(w http.ResponseWriter, r *http.Request, status int) {
	RespondErr(w, r, status, http.StatusText(status))
}
