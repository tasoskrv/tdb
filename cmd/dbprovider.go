package main

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBconn struct {
	client *mongo.Client
	db     *mongo.Database
}

//InitDatabase fns
func InitDatabase() (DBconn, func(), error) {
	port := "2717"

	if len(os.Getenv("mport")) != 0 {
		port = os.Getenv("mport")
	}

	connMongo := "mongodb://127.0.0.1:" + port
	dbMongo := "tdb"

	//Create client
	client, err := mongo.NewClient(options.Client().ApplyURI(connMongo))
	if err != nil {
		return DBconn{}, nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	//Connect to mongo
	err = client.Connect(ctx)
	if err != nil {
		return DBconn{}, nil, err
	}

	// Return disconnect as func to defer in main
	f := func() {
		client.Disconnect(ctx)
	}

	//Ping Database
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return DBconn{}, f, err
	}
	db := client.Database(dbMongo)

	return DBconn{client: client, db: db}, f, nil

}
