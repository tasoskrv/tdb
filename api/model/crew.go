package model

import (
	"encoding/json"
	"net/http"
)

//Crew structure
type Crew struct {
	Tconst    string `bson:"tconst" json:"tconst"`       //alphanumeric unique identifier of the title
	Directors string `bson:"directors" json:"directors"` //(string array)
	Writers   string `bson:"writers" json:"writers"`     //(string array)
}

/*
//MongoCon structure
type MongoCon struct {
	Client   *mongo.Client
	Database *mongo.Database
}
*/
/*
type response struct {
	success bool
	id      string
	data    Movie
}

type responseDelete struct {
	Success bool
}*/

//CreateMovie ... creates a movie
func (mongocon *MongoCon) CreateCrew(w http.ResponseWriter, r *http.Request) {
	//setHeaders(w)
	/*
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
	*/
	res := ResponseDelete{
		Success: true,
	}

	b, _ := json.Marshal(res)
	w.Write(b)
}

//UpdateMovie updates movie data
func (mongocon *MongoCon) UpdateCrew(w http.ResponseWriter, r *http.Request) {
	//setHeaders(w)
	/*
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
	*/
}

//DeleteMovie deletes a movie
func (mongocon *MongoCon) DeleteCrew(w http.ResponseWriter, r *http.Request) {
	//setHeaders(w)
	/*
		params := mux.Vars(r) //Get params

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		moviesCollection := mongocon.Database.Collection("movie")

		_, err := moviesCollection.DeleteOne(ctx, bson.M{"tconst": params["tconst"]})
		if err != nil {
			writeErrorLogs(err)
		}

		res := responseDelete{
			Success: true,
		}

		b, _ := json.Marshal(res)

		w.Write(b)
		//writeResponse(w, b, errR)
	*/
}

/*
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
}*/
