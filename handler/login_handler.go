package handler

import (
	"../model"
	"github.com/labstack/echo"
	"net/http"
)

type LoginParam struct {
	Email string
	Password string
}

func (h *Handler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(LoginParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		user := model.User{}
		h.DB.First(&user, "email=? and password=?", param.Email, param.Password)
		if h.DB.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]error{
				"error": h.DB.Error,
			})
		}

		t, err := user.IssueToken()
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]error{
				"error": err,
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
}

