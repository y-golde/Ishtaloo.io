package rooms

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"ishtaloo.io/DB/Collections"
	entities "ishtaloo.io/Entities"
)

func getRooms(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	roomsCollection := Collections.RoomsGetCollection()
	cursor, err := roomsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var rooms []entities.Room
	if err = cursor.All(ctx, &rooms); err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, rooms)
}
