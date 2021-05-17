package rootController

import (
	"github.com/labstack/echo/v4"
	"ishtaloo.io/API/auth"
	"ishtaloo.io/API/index"
	"ishtaloo.io/API/rooms"
	"ishtaloo.io/API/users"
	"ishtaloo.io/API/words"
	entities "ishtaloo.io/Entities"
)

func RootController(e *echo.Echo, sseChanel *entities.SSEChannel) {
	users.UsersController(e)
	index.IndexController(e)
	words.WordsController(e)
	auth.AuthController(e)
	rooms.RoomsController(e, sseChanel)
}
