package leave

import (
	"github.com/labstack/echo/v4"
	entities "ishtaloo.io/Entities"
)

func LeaveController(e *echo.Echo, sseChannel *entities.SSEChannel) {
	e.POST("/rooms/leave/:roomId", func(c echo.Context) error { return postLeave(c, sseChannel) })
}
