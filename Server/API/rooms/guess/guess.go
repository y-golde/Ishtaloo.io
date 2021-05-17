package guess

import (
	"github.com/labstack/echo/v4"
)

func GuessController(e *echo.Echo) {
	e.POST("/rooms/guess", postGuess)
}
