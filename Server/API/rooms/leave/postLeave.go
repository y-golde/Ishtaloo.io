package leave

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

func postLeave(c echo.Context, sseChannel *entities.SSEChannel) error {
	roomId := c.Param("roomId")
	u, _ := cookieUtils.GetUserFromCookie(c)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := roomsScripts.RemoveUserFromRoom(ctx, roomId, u); err != nil {
		return c.String(http.StatusInternalServerError, strings.Join([]string{err.Error(), "from RemoveUserFromRoom"}, ", "))
	}

	sseChannel.Notifier <- "users changed"

	return c.NoContent(http.StatusOK)
}
