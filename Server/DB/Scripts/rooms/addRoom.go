package roomsScripts

import (
	"context"
	"fmt"
	"log"

	"ishtaloo.io/DB/Collections"
	entities "ishtaloo.io/Entities"
)

func AddRoom(ctx context.Context, room entities.Room) {
	roomsCollection := Collections.RoomsGetCollection()
	newRoomsResult, err := roomsCollection.InsertOne(ctx, room)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted single document: ", newRoomsResult.InsertedID)
}
