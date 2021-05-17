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
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	u, _ := cookieUtils.GetUserFromCookie(c)
	userInRoom, err := roomsScripts.UserInRoom(ctx, roomId, u)
	if err != nil {
		return c.String(http.StatusInternalServerError, strings.Join([]string{err.Error(), "from UserInRoom"}, ", "))
	}
	if !userInRoom {
		return c.String(http.StatusForbidden, "User not in room.")
	}

	g := new(struct{ Guess string })
	if err = c.Bind(g); err != nil {
		return c.String(http.StatusForbidden, "body doesn't contain string.")
	}

	if len(g.Guess) != 1 {
		return c.String(http.StatusForbidden, "body contains more than single character.")
	}
	r := rune(g.Guess[0])

	if err = roomsScripts.AddGuessToRoom(ctx, roomId, r); err != nil {
		return c.String(http.StatusInternalServerError, strings.Join([]string{err.Error(), "from AddGuessToRoom"}, ", "))
	}
	return c.NoContent(http.StatusOK)
}
