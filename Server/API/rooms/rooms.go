package rooms

import (
	"github.com/labstack/echo/v4"
)

func RoomsController(e *echo.Echo) {
	e.POST("/rooms", postRoom)
}
