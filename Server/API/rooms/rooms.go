package rooms

import (
	"github.com/labstack/echo/v4"
	"ishtaloo.io/API/rooms/guess"
	"ishtaloo.io/API/rooms/join"
	"ishtaloo.io/API/rooms/leave"
	roomUsers "ishtaloo.io/API/rooms/users"
	entities "ishtaloo.io/Entities"
)

func RoomsController(e *echo.Echo, sseChannel *entities.SSEChannel) {
	e.POST("/rooms", postRoom)
	e.GET("/rooms", getRooms)
	e.GET("/rooms/:roomId", getRoomById)

	join.JoinController(e, sseChannel)
	leave.LeaveController(e, sseChannel)
	guess.GuessController(e, sseChannel)
	roomUsers.UsersController(e)
}
