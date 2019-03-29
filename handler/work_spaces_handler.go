package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"../model"
)

type WorkSpaceParam struct {
	Name string
}

func (h *Handler) GetWorkSpace() echo.HandlerFunc {
	return func(c echo.Context) error {
		workSpaceId := c.Param("id")
		workSpace := model.WorkSpace{}
		result := h.DB.Preload("Users").First(&workSpace, "id=?", workSpaceId)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"status": "Not Found",
			})
		}
		return c.JSON(http.StatusOK, struct {
			WorkSpace model.WorkSpace `json:"work_space"`
		} {
			WorkSpace: workSpace,
		})
	}
}

func (h *Handler) SaveWorkSpace() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(WorkSpaceParam)
		if err := c.Bind(param); err != nil {
			return err
		}

		workSpace := model.WorkSpace{
			Name: param.Name,
		}
		h.DB.Create(&workSpace)

		userName := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["name"].(string)
		user := model.User{}
		h.DB.First(&user, "name=?", userName)
		h.DB.Model(&workSpace).Association("Users").Append(&user)
		return c.JSON(http.StatusOK, echo.Map{
			"Name": workSpace.Name,
		})
	}
}
