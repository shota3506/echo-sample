package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"../model"
)

type teamMemberParam struct {
	Name string `json:"name"`
}

func (h *Handler) GetTeamMembers() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("team_id")
		team := model.Team{}
		h.DB.Preload("Members").First(&team, "id=?", teamId)
		if h.DB.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]error{
				"error": h.DB.Error,
			})
		}
		return c.JSON(http.StatusOK, struct {
			Members []model.Member `json:"members"`
		} {
			Members: team.Members,
		})
	}
}

func (h *Handler) CreateTeamMember() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("team_id")
		team := model.Team{}
		h.DB.Preload("Members").First(&team, "id=?", teamId)
		if h.DB.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]error{
				"error": h.DB.Error,
			})
		}

		param := new(teamMemberParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		userEmail := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
		user := model.User{}
		result := h.DB.First(&user, "email=?", userEmail)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]error{
				"error": result.Error,
			})
		}
		member := model.Member{
			User: user,
			Team: team,
			Name: param.Name,
			Role: "general",
		}
		if err := c.Validate(member); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":  err.Error(),
			})
		}
		result = h.DB.Create(&member)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]error{
				"error": result.Error,
			})
		}
		return c.JSON(http.StatusOK, struct {
			Member model.Member `json:"member"`
		} {
			Member: member,
		})
	}
}