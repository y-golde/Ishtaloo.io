package login

import (
	"github.com/labstack/echo/v4"
)

func LoginController(e *echo.Echo) {
	e.POST("/login", postLogin)
}
