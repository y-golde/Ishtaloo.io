package guess

import (
	"github.com/labstack/echo/v4"
	entities "ishtaloo.io/Entities"
)

func GuessController(e *echo.Echo, sseChannel *entities.SSEChannel) {
	e.POST("/rooms/guess", func(c echo.Context) error { return postGuess(c, sseChannel) })
}
