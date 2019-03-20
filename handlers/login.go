package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/shota3506/echo_sample/models"
	"net/http"
	"time"
)

type LoginParam struct {
	Name string
	Password string
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := gorm.Open("mysql", "root:@/echo_sample?parseTime=true")
		if err != nil {
			panic("データベースへの接続に失敗しました")
		}
		defer db.Close()

		param := new(LoginParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		user := model.User{}
		result := db.First(&user, "name=? and password=?", param.Name, param.Password)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"status": "Not Fount",
			})
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = user.Name
		claims["admin"] = false
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("wkGRdkcF2taUE"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
}

