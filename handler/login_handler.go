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
		if err := c.Bind(param); err != nil { return h.return400(c, err) }
		user := model.User{}
		result := h.DB.First(&user, "email=? and password=?", param.Email, param.Password)
		if result.Error != nil { return h.return404(c, result.Error) }

		t, _ := user.IssueToken()
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
}

