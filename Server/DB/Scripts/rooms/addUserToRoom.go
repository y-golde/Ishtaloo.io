package roomsScripts

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ishtaloo.io/DB/Collections"
	entities "ishtaloo.io/Entities"
)

func AddUserToRoom(ctx context.Context, roomId string, user *entities.User) {
	roomsCollection := Collections.RoomsGetCollection()
	id, _ := primitive.ObjectIDFromHex(roomId)
	var room entities.Room
	if err := roomsCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&room); err != nil {
		log.Fatal(err)
	}
	users := room.Users

	if !userAlreadyExists(user,users) {
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

func userAlreadyExists(user *entities.User, users []entities.User) bool {
	for _, curUser := range users {
		if(curUser.UserId == user.UserId) {
			return true
		}
	}
	
	return false
}
