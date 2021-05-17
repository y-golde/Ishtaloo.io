package auth

import (
	"github.com/labstack/echo/v4"
)

func AuthController(e *echo.Echo) {
	e.POST("/auth", postAuth)
}
