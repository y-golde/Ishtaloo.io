package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id    string `json:"id" bson:"id"`
	Email string `json:"email" bson:"email"`
}

func getUser(c echo.Context) error {
	u := &User{
		Id:    "1",
		Email: "Dani@gmail.com",
	}
	return c.JSON(http.StatusOK, u)
}
