package guess

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	roomsScripts "ishtaloo.io/DB/Scripts/rooms"
	entities "ishtaloo.io/Entities"
	cookieUtils "ishtaloo.io/Utils/cookie"
)

func postGuess(c echo.Context, sseChannel *entities.SSEChannel) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var g entities.Guess
	if err = c.Bind(&g); err != nil {
		return c.String(http.StatusForbidden, "body bad format.")
	}

	u, _ := cookieUtils.GetUserFromCookie(c)
	userInRoom, err := roomsScripts.UserInRoom(ctx, g.RoomId, u)
	if err != nil {
		return c.String(http.StatusInternalServerError, strings.Join([]string{err.Error(), "from UserInRoom"}, ", "))
	}
	if !userInRoom {
		return c.String(http.StatusForbidden, "User not in room.")
	}

	if len(g.Guess) != 1 {
		return c.String(http.StatusForbidden, "body contains more than single character.")
	}
	r := rune(g.Guess[0])

	if err = roomsScripts.AddGuessToRoom(ctx, g.RoomId, r); err != nil {
		return c.String(http.StatusInternalServerError, strings.Join([]string{err.Error(), "from AddGuessToRoom"}, ", "))
	}

	sseChannel.Notifier <- "new guess"
	return c.NoContent(http.StatusOK)
}
