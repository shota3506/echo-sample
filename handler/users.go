package handler

import (
	"net/http"
	"time"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"../model"

)

type UserParam struct {
	Name string
	Password string
}

func (h *Handler) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("id")
		user := model.User{}
		result := h.DB.First(&user, "id=?", userId)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"status": "Not Found",
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"Name": user.Name,
		})
	}
}

func (h *Handler) SaveUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(UserParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		user := model.User{}
		user.Name = param.Name
		user.Password = param.Password
		h.DB.Create(&user)

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
