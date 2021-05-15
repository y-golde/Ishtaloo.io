package wordsScripts

import (
	"context"
	"fmt"
	"log"

	"ishtaloo.io/DB/Collections"
	entities "ishtaloo.io/Entities"
)

func AddWord(ctx context.Context, word entities.Word) {
	wordsCollection := Collections.WordsGetCollection()
	newWordsResult, err := wordsCollection.InsertOne(ctx, word)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted single document: ", newWordsResult.InsertedID)
}
