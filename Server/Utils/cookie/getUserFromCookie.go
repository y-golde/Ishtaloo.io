package cookieUtils

import (
	"fmt"

	"github.com/labstack/echo/v4"
	entities "ishtaloo.io/Entities"
)

func GetUserFromCookie(c echo.Context) (*entities.User, error) {
	//user := new(entities.User)
	cookie, err := c.Cookie("user")
	if err != nil {
		return nil, err
	}
	u, _ := decodedToken(cookie.Value)
	fmt.Println(u)

	user := &entities.User{
		UserName: u["user_name"].(string),
		UserId:   u["user_id"].(string),
	}
	return user, nil
}
