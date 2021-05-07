package words

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"ishtaloo.io/DB/Collections"
)

func getWords(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	wordsCollection := Collections.WordsGetCollection()
	cursor, err := wordsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var words []bson.M
	if err = cursor.All(ctx, &words); err != nil {
		log.Fatal(err)
	}
	fmt.Println(words)
	return c.JSON(http.StatusOK, words)
}
