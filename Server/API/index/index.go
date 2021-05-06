package index

import (
	"github.com/labstack/echo/v4"
)

func IndexController(e *echo.Echo) {
	e.GET("/", getIndex)
}
