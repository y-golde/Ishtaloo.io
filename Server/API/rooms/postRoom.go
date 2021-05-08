package rooms

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"ishtaloo.io/DB/Scripts"
	entities "ishtaloo.io/Entities"
	Util "ishtaloo.io/Utils"
)

func postRoom(c echo.Context) (err error) {
	//r := new(map[string]interface{})
	r := new(entities.Room)
	if err = c.Bind(r); err != nil {
		return err
	}
	u, _ := Util.GetUserFromCookie(c)
	users := make([]entities.User, 10)
	users = append([]entities.User{*u}, users...)
	room := entities.Room{
		RoomId:      r.RoomId,
		CurrentWord: r.CurrentWord,
		Guesses:     r.Guesses,
		Users:       users,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Scripts.AddRoom(ctx, room)

	return c.NoContent(http.StatusCreated)
}
