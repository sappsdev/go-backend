package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CodesCollection *mongo.Collection
var UserCollection *mongo.Collection


func Collections() {
	CodesCollection = DB.Collection("codes")
	UserCollection = DB.Collection("users")

	_, _ = UserCollection.Indexes().CreateOne(Ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})

	_, _ = CodesCollection.Indexes().CreateOne(Ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "created_at", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(300),
	})

}
