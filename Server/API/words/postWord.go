package words

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	wordsScripts "ishtaloo.io/DB/Scripts/words"
	entities "ishtaloo.io/Entities"
)

func postWord(c echo.Context) (err error) {
	w := new(struct{ Word string })
	if err = c.Bind(w); err != nil {
		return err
	}
	for _, r := range w.Word {
		if !runeValid(r) {
			return c.String(http.StatusForbidden, "Word string not valid.")
		}
	}
	word := entities.Word{
		Word: []rune(w.Word),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wordsScripts.AddWord(ctx, word)

	return c.String(http.StatusOK, "word wrod")
}

func runeValid(r rune) bool {
	if r >= 65 && r <= 90 || r == 32 || r == 45 {
		return true
	}
	return false
}
