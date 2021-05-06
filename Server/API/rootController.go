package rootController

import (
	"github.com/labstack/echo/v4"
	"ishtaloo.io/API/index"
	"ishtaloo.io/API/users"
)

func RootController(e *echo.Echo) {
	users.UsersController(e)
	index.IndexController(e)
}
