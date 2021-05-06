package words

import (
	"github.com/labstack/echo/v4"
)

func WordsController(e *echo.Echo) {
	e.GET("/words", getWords)
	e.POST("/words", postWord)
}
