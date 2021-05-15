package wordsScripts

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"ishtaloo.io/DB/Collections"
	entities "ishtaloo.io/Entities"
)

func GetWordsByFilter(filter primitive.M) []entities.Word {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	wordsCollection := Collections.WordsGetCollection()
	cursor, err := wordsCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	var words []entities.Word
	if err = cursor.All(ctx, &words); err != nil {
		log.Fatal(err)
	}
	fmt.Println(words)
	return words
}
