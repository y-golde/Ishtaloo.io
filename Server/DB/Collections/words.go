package Collections

import (
	"go.mongodb.org/mongo-driver/mongo"
	"ishtaloo.io/DB"
)

func WordsGetCollection() *mongo.Collection {
	return DB.Client.Database("ishtaloo").Collection("words")
}
