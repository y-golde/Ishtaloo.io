package index

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getIndex(c echo.Context) error {
	fmt.Printf("wowowowo")
	return c.String(http.StatusOK, "Hello, World!\n")
}
