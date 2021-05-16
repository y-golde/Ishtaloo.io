package roomsScripts

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ishtaloo.io/DB/Collections"
	entities "ishtaloo.io/Entities"
)

func GetUsersOfRoom(ctx context.Context, roomId string) ([]entities.User, error) {
	roomsCollection := Collections.RoomsGetCollection()
	id, _ := primitive.ObjectIDFromHex(roomId)
	var room entities.Room
	if err := roomsCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&room); err != nil {
		return nil, err
	}
	users := room.Users
	return users, nil
}
