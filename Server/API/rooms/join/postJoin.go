package join

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	roomsScripts "ishtaloo.io/DB/Scripts/rooms"
	entities "ishtaloo.io/Entities"
	cookieUtils "ishtaloo.io/Utils/cookie"
)

func postJoin(c echo.Context, sseChannel *entities.SSEChannel) (err error) {
	roomId := c.Param("roomId")
	u, _ := cookieUtils.GetUserFromCookie(c)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	roomsScripts.AddUserToRoom(ctx, roomId, u)

	sseChannel.Notifier <- "users changed"

	return c.NoContent(http.StatusOK)
}
