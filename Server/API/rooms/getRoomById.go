package rooms

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ishtaloo.io/DB/Collections"
	entities "ishtaloo.io/Entities"
	roomUtils "ishtaloo.io/Utils/room"
)

func getRoomById(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	roomsCollection := Collections.RoomsGetCollection()
	roomId := c.Param("roomId")
	id, _ := primitive.ObjectIDFromHex(roomId)
	var room entities.Room
	if err := roomsCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&room); err != nil {
		log.Fatal(err)
	}
	roomUtils.EncryptWord(room.CurrentWord, room.Guesses)

	clientRoom := struct {
		RoomId      string
		CurrentWord string
		Guesses     []rune
		Users       []entities.User
	}{
		RoomId:      room.RoomId,
		CurrentWord: string(room.CurrentWord),
		Guesses:     room.Guesses,
		Users:       room.Users,
	}

	return c.JSON(http.StatusOK, clientRoom)
}
