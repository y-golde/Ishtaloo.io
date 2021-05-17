package roomsScripts

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ishtaloo.io/DB/Collections"
	entities "ishtaloo.io/Entities"
	wordUtils "ishtaloo.io/Utils/word"
)

func AddGuessToRoom(ctx context.Context, roomId string, guess rune) error {
	if guessValid(guess) {
		return errors.New("guess is not a capital English letter")
	}

	roomsCollection := Collections.RoomsGetCollection()
	id, _ := primitive.ObjectIDFromHex(roomId)
	var room entities.Room
	if err := roomsCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&room); err != nil {
		return err
	}

	takenGuesses := append(room.Guesses, wordUtils.VisibleRunes()...)
	for _, g := range takenGuesses {
		if g == guess {
			return errors.New("guess already exists")
		}
	}
	room.Guesses = append(room.Guesses, guess)

	_, err := roomsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			primitive.E{
				Key: "$set",
				Value: bson.D{primitive.E{
					Key:   "guesses",
					Value: room.Guesses,
				}},
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func guessValid(g rune) bool {
	if g < 65 || g > 90 {
		return false
	}
	return true
}
