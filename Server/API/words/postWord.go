package words

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"ishtaloo.io/DB/Scripts"
	entities "ishtaloo.io/Entities"
)

func postWord(c echo.Context) (err error) {
	w := new(entities.Word)
	if err = c.Bind(w); err != nil {
		return err
	}
	word := entities.Word{
		Word: w.Word,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Scripts.AddWord(ctx, word)

	return c.String(http.StatusOK, "word wrod")
}
