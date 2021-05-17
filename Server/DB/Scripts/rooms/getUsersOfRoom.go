package roomsScripts

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ishtaloo.io/DB/Collections"
	entities "ishtaloo.io/Entities"
)

func GetUsersOfRoom(ctx context.Context, roomId string) ([]entities.User, error) {
	roomsCollection := Collections.RoomsGetCollection()
	id, _ := primitive.ObjectIDFromHex(roomId)
	var users struct{ Users []entities.User }
	if err := roomsCollection.FindOne(
		ctx,
		bson.M{"_id": id},
		options.FindOne().SetProjection(bson.M{"users": 1, "_id": 0}),
	).Decode(&users); err != nil {
		return nil, err
	}
	return users.Users, nil
}
