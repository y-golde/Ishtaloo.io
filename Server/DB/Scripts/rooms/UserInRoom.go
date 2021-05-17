package roomsScripts

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ishtaloo.io/DB/Collections"
	entities "ishtaloo.io/Entities"
)

func UserInRoom(ctx context.Context, roomId string, user *entities.User) (bool, error) {
	roomsCollection := Collections.RoomsGetCollection()
	id, _ := primitive.ObjectIDFromHex(roomId)
	var room entities.Room
	if err := roomsCollection.FindOne(ctx, bson.M{"_id": id, "users": primitive.M{"$elemMatch": primitive.M{"user_id": user.UserId}}}).Decode(&room); err != nil {
		return false, err
	}

	return true, nil

}
