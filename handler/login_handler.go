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
		result := h.DB.First(&user, "email=? and password=?", param.Email, param.Password)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"status": "Not Found",
			})
		}

		t, err := user.IssueToken()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
}

