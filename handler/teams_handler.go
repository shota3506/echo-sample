package handler

import (
	"../model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
)

type teamParam struct {
	Name string `json:"name"`
	MemberName string `json:"member_name"`
}

func (h *Handler) GetTeams() echo.HandlerFunc {
	return func(c echo.Context) error {
		userEmail := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
		user := model.User{}
		result := h.DB.Preload("Teams").First(&user, "email=?", userEmail)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]error{
				"error": result.Error,
			})
		}
		return c.JSON(http.StatusOK, struct {
			Teams []model.Team `json:"teams"`
		} {
			Teams: user.Teams,
		})
	}
}

func (h *Handler) GetTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("id")
		team := model.Team{}
		h.DB.Preload("Users").First(&team, "id=?", teamId)
		if h.DB.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]error{
				"error": h.DB.Error,
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
		param := new(teamParam)
		if err := c.Bind(param); err != nil {
			return err
		}

		//rootFolder := model.Folder{Title: "root"}
		//h.DB.Create(&rootFolder)
		team := model.Team{
			Name: param.Name,
			RootFolder: model.Folder{Title: "root"},
		}
		result := h.DB.Create(&team)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]error{
				"error": result.Error,
			})
		}
		userEmail := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
		user := model.User{}
		result = h.DB.First(&user, "email=?", userEmail)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]error{
				"error": result.Error,
			})
		}

		member := model.Member{
			User: user,
			Team: team,
			Name: param.MemberName,
			Role: "admin",
		}
		result = h.DB.Create(&member)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]error{
				"error": result.Error,
			})
		}
		return c.JSON(http.StatusOK, struct {
			Team model.Team `json:"team"`
		} {
			Team: team,
		})
	}
}
