package roomUsers

import (
	"github.com/labstack/echo/v4"
)

func UsersController(e *echo.Echo) {
	e.GET("/rooms/:roomId/users", getUsers)
}
