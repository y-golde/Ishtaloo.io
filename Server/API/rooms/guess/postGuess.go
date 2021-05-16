package guess

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	roomsScripts "ishtaloo.io/DB/Scripts/rooms"
	cookieUtils "ishtaloo.io/Utils/cookie"
)

func postGuess(c echo.Context) (err error) {
	roomId := c.Param("roomId")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	u, _ := cookieUtils.GetUserFromCookie(c)
	usersInRoom, err := roomsScripts.GetUsersOfRoom(ctx, roomId)
	if err != nil {
		return c.String(http.StatusInternalServerError, strings.Join([]string{err.Error(), "from GetUserOfRoom"}, ", "))
	}
	userInRoom := false
	for _, v := range usersInRoom {
		if u.UserId == v.UserId {
			userInRoom = true
			break
		}
	}
	if !userInRoom {
		return c.String(http.StatusForbidden, "User not in room.")
	}

	g := new(struct{ Guess string })
	if err = c.Bind(g); err != nil {
		return c.String(http.StatusForbidden, "body doesn't contain string.")
	}
	var r rune
	for _, v := range g.Guess {
		if r != 0 {
			return c.String(http.StatusForbidden, "body contains more than single character.")
		}
		r = v
	}

	if err = roomsScripts.AddGuessToRoom(ctx, roomId, r); err != nil {
		return c.String(http.StatusInternalServerError, strings.Join([]string{err.Error(), "from AddGuessToRoom"}, ", "))
	}
	return c.NoContent(http.StatusOK)
}
