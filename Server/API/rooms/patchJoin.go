package rooms

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	roomsScripts "ishtaloo.io/DB/Scripts/rooms"
	cookieUtils "ishtaloo.io/Utils/cookie"
)

func patchJoin(c echo.Context) (err error) {
	roomId := c.Param("roomId")
	u, _ := cookieUtils.GetUserFromCookie(c)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	roomsScripts.AddUserToRoom(ctx, roomId, u)

	return c.NoContent(http.StatusOK)
}
