package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"../model"
)

type teamSpaceParam struct {
	Name string
}

func (h *Handler) GetTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("id")
		team := model.Team{}
		result := h.DB.Preload("Users").First(&team, "id=?", teamId)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"status": "Not Found",
			})
		}
		return c.JSON(http.StatusOK, struct {
			Team model.Team `json:"team"`
		} {
			Team: team,
		})
	}
}

func (h *Handler) CreateTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(teamSpaceParam)
		if err := c.Bind(param); err != nil {
			return err
		}

		team := model.Team{
			Name: param.Name,
		}
		h.DB.Create(&team)

		userEmail := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
		user := model.User{}
		h.DB.First(&user, "email=?", userEmail)
		h.DB.Model(&team).Association("Users").Append(&user)
		return c.JSON(http.StatusOK, echo.Map{
			"Name": team.Name,
		})
	}
}
