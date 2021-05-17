package rooms

import (
	"github.com/labstack/echo/v4"
	"ishtaloo.io/API/rooms/guess"
	"ishtaloo.io/API/rooms/join"
	roomUsers "ishtaloo.io/API/rooms/users"
)

func RoomsController(e *echo.Echo) {
	e.POST("/rooms", postRoom)
	e.GET("/rooms", getRooms)
	e.GET("/rooms/:roomId", getRoomById)

	join.JoinController(e)
	guess.GuessController(e)
	roomUsers.UsersController(e)
}
