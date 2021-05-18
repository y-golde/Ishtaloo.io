package sse

import (
	"github.com/labstack/echo/v4"
	entities "ishtaloo.io/Entities"
)

func SSEController(e *echo.Echo, sseChannel *entities.SSEChannel) {
	e.GET("/sse", func(c echo.Context) error { return getSSE(c, sseChannel) })
}
