package join

import (
	"github.com/labstack/echo/v4"
)

func JoinController(e *echo.Echo) {
	e.POST("/rooms/join/:roomId", postJoin)
}
