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
		result := h.DB.Preload("Members").First(&team, "id=?", teamId)
		if result.Error != nil { return h.return404(c, result.Error) }
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
		result := h.DB.Preload("Members").First(&team, "id=?", teamId)
		if result.Error != nil { return h.return404(c, result.Error) }

		param := new(teamMemberParam)
		if err := c.Bind(param); err != nil { return h.return400(c, err) }
		userEmail := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
		user := model.User{}
		result = h.DB.First(&user, "email=?", userEmail)
		if result.Error != nil { return h.return404(c, result.Error) }
		member := model.Member{
			User: user,
			Team: team,
			Name: param.Name,
			Role: "general",
		}
		result = h.DB.Create(&member)
		if result.Error != nil { return h.return400(c, result.Error) }
		return c.JSON(http.StatusOK, struct {
			Member model.Member `json:"member"`
		} {
			Member: member,
		})
	}
}