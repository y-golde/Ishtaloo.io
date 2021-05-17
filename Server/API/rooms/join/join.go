package join

import (
	"github.com/labstack/echo/v4"
	entities "ishtaloo.io/Entities"
)

func JoinController(e *echo.Echo, sseChannel *entities.SSEChannel) {
	e.POST("/rooms/join/:roomId", func(c echo.Context) error { return postJoin(c, sseChannel) })
}
