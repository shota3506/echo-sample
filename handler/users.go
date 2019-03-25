package handler

import (
	"net/http"
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"../model"
)

type UserParam struct {
	Name string
	Password string
}

func GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := gorm.Open("mysql", os.Getenv("DATABASE_SOURCE"))
		if err != nil {
			panic("データベースへの接続に失敗しました")
		}
		defer db.Close()

		userId := c.Param("id")
		user := model.User{}
		result := db.First(&user, "id=?", userId)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"status": "Not Fount",
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"Name": user.Name,
		})
	}
}

func SaveUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := gorm.Open("mysql", os.Getenv("DATABASE_SOURCE"))
		if err != nil {
			panic("データベースへの接続に失敗しました")
		}
		defer db.Close()

		param := new(UserParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		user := model.User{}
		user.Name = param.Name
		user.Password = param.Password
		db.Create(&user)

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = user.Name
		claims["admin"] = false
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Create token with claims

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("wkGRdkcF2taUE"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
}
