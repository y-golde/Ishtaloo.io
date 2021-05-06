package words

import (
	"github.com/labstack/echo/v4"
)

func WordsController(e *echo.Echo) {
	e.POST("/words", postWord)
}
