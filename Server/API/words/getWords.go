package words

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	wordsScripts "ishtaloo.io/DB/Scripts/words"
)

func getWords(c echo.Context) error {
	words := wordsScripts.GetWordsByFilter(bson.M{})
	return c.JSON(http.StatusOK, words)
}
