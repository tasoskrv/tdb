package model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

//MongoCon structure
type MongoCon struct {
	Client   *mongo.Client
	Database *mongo.Database
}
