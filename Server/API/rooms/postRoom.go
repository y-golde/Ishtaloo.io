package rooms

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	roomsScripts "ishtaloo.io/DB/Scripts/rooms"
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
	room := entities.Room{
		CurrentWord: r.CurrentWord,
		Guesses:     r.Guesses,
		Users:       users,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	roomsScripts.AddRoom(ctx, room)

	return c.NoContent(http.StatusCreated)
}
