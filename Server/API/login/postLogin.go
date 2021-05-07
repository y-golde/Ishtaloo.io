package login

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func postLogin(c echo.Context) (err error) {
	cookie := new(http.Cookie)
	cookie.Name = "userId"
	cookie.Value = createToken()
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "login successful")
}

func createToken() string {
	s := "hello"
	return s
}
