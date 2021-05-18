package roomsScripts

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ishtaloo.io/DB/Collections"
	entities "ishtaloo.io/Entities"
	roomUtils "ishtaloo.io/Utils/room"
)

func AddUserToRoom(ctx context.Context, roomId string, user *entities.User) {
	roomsCollection := Collections.RoomsGetCollection()
	id, _ := primitive.ObjectIDFromHex(roomId)
	var room entities.Room
	if err := roomsCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&room); err != nil {
		log.Fatal(err)
	}
	users := room.Users

	if roomUtils.UserExistsInSlice(user, users) == -1 {
		users = append([]entities.User{*user}, users...)

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
			log.Fatal(err)
		}
	}
}
