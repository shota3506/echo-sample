package handler

import (
	"../model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"net/http"
)

type userParam struct {
	Email string
	Password string
}

func (h *Handler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("id")
		user := model.User{}
		result := h.DB.Preload("Teams").First(&user, "id=?", userId)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]error{
				"error": result.Error,
			})
		}
		return c.JSON(http.StatusOK, struct {
			User model.User `json:"user"`
		} {
			User: user,
		})
	}
}

func (h *Handler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(userParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		user := model.User{
			Email: param.Email,
			Password: param.Password,
		}
		result := h.DB.Create(&user)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]error{
				"error": result.Error,
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
