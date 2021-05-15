package rooms

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	roomsScripts "ishtaloo.io/DB/Scripts/rooms"
	wordsScripts "ishtaloo.io/DB/Scripts/words"
	entities "ishtaloo.io/Entities"
	cookieUtils "ishtaloo.io/Utils/cookie"
)

func postRoom(c echo.Context) (err error) {
	r := new(entities.Room)
	if err = c.Bind(r); err != nil {
		return err
	}

	u, _ := cookieUtils.GetUserFromCookie(c)
	const maxUsers int = 10
	users := make([]entities.User, maxUsers)
	users = append([]entities.User{*u}, users...)

	words := wordsScripts.GetWordsByFilter(bson.M{})
	randomIndex := rand.Intn(len(words))

	guesses := *new([]rune)
	room := entities.Room{
		CurrentWord: words[randomIndex].Word,
		Guesses:     guesses,
		Users:       users,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	roomsScripts.AddRoom(ctx, room)

	return c.NoContent(http.StatusCreated)
}
