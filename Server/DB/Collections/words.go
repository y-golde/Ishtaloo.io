package Collections

import (
	"go.mongodb.org/mongo-driver/mongo"
	"ishtaloo.io/DB"
)

func WordsGetCollection() *mongo.Collection {
	return DB.Client.Database("router-test").Collection("words")
}
