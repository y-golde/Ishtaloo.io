package rootController

import (
	"github.com/labstack/echo/v4"
	"ishtaloo.io/API/index"
	"ishtaloo.io/API/login"
	"ishtaloo.io/API/users"
	"ishtaloo.io/API/words"
)

func RootController(e *echo.Echo) {
	users.UsersController(e)
	index.IndexController(e)
	words.WordsController(e)
	login.LoginController(e)
}
