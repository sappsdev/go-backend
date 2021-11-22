package database

import (
	"context"
	"time"

	"backend/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx context.Context

var Cancel context.CancelFunc

var Client *mongo.Client

var DB *mongo.Database

func Connect() {
	var err error

	Ctx, Cancel = context.WithTimeout(context.Background(), 30*time.Second)
	Client, err = mongo.Connect(Ctx, options.Client().ApplyURI(config.MongoUri()))
	if err != nil {
		panic(err)
	}
	DB = Client.Database(config.Database())
}

func IsDupKey(err error) bool {
	if wes, ok := err.(mongo.WriteException); ok {
		for i := range wes.WriteErrors {
			if wes.WriteErrors[i].Code == 11000 || wes.WriteErrors[i].Code == 11001 || wes.WriteErrors[i].Code == 12582 || wes.WriteErrors[i].Code == 16460 {
				return true
			}
		}
	}
	return false
}
