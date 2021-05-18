package roomsScripts

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ishtaloo.io/DB/Collections"
	entities "ishtaloo.io/Entities"
	roomUtils "ishtaloo.io/Utils/room"
)

func RemoveUserFromRoom(ctx context.Context, roomId string, user *entities.User) error {
	roomsCollection := Collections.RoomsGetCollection()
	id, _ := primitive.ObjectIDFromHex(roomId)
	var room entities.Room
	if err := roomsCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&room); err != nil {
		log.Fatal(err)
	}
	users := room.Users

	if u := roomUtils.UserExistsInSlice(user, users); u > -1 {
		users = append(users[:u], users[u+1:]...)

		_, err := roomsCollection.UpdateOne(
			ctx,
			bson.M{"_id": id},
			bson.D{
				primitive.E{
					Key: "$set",
					Value: bson.D{primitive.E{
						Key:   "users",
						Value: users,
					}}},
			},
		)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("user with id: %v, doesn't exist in room: %v", user.UserId, roomId)
}
