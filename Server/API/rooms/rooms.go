package rooms

import (
	"github.com/labstack/echo/v4"
)

func RoomsController(e *echo.Echo) {
	e.POST("/rooms", postRoom)
	e.PATCH("/rooms/join/:roomId", patchJoin)
	e.GET("/rooms", getRooms)
	e.GET("/rooms/:roomId", getRoomById)
}
