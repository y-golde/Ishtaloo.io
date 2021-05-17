package roomUsers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	roomsScripts "ishtaloo.io/DB/Scripts/rooms"
)

func getUsers(c echo.Context) (err error) {
	roomId := c.Param("roomId")

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	users, err := roomsScripts.GetUsersOfRoom(ctx, roomId)
	if err != nil {
		return c.String(http.StatusInternalServerError, strings.Join([]string{err.Error(), "from UserInRoom"}, ", "))
	}

	return c.JSON(http.StatusOK, users)
}
