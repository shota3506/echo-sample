package handler

import (
	"../model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type LoginParam struct {
	Name string
	Password string
}

func (h *Handler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(LoginParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		user := model.User{}
		result := h.DB.First(&user, "name=? and password=?", param.Name, param.Password)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"status": "Not Found",
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

