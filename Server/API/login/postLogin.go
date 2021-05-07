package login

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	entities "ishtaloo.io/Entities"
)

func postLogin(c echo.Context) (err error) {
	cookie := new(http.Cookie)
	cookie.Name = "userId"
	u := new(entities.LoginRequest)
	if err = c.Bind(u); err != nil {
		return err
	}
	userName := entities.LoginRequest{
		UserName: u.UserName,
	}
	cookie.Value, err = createToken(userName)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	cookie.HttpOnly = true
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "login successful")
}

func createToken(userName entities.LoginRequest) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_name"] = userName.UserName
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
