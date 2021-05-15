package join

import (
	"github.com/labstack/echo/v4"
)

func JoinController(e *echo.Echo) {
	e.PATCH("/rooms/join/:roomId", patchJoin)
}
