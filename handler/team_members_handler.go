package handler

import (
	"../model"
	"github.com/labstack/echo"
	"net/http"
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
		e := h.setCurrentUser(c)
		if e != nil { return h.return404(c, e) }

		teamId := c.Param("team_id")
		team := model.Team{}
		result := h.DB.Preload("Members").First(&team, "id=?", teamId)
		if result.Error != nil { return h.return404(c, result.Error) }

		param := new(teamMemberParam)
		if err := c.Bind(param); err != nil { return h.return400(c, err) }

		member := model.Member{
			User: h.CurrentUser,
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
		if result.Error != nil { return h.return400(c, result.Error) }
		return c.JSON(http.StatusOK, struct {
			Member model.MemberResponse `json:"member"`
		} {
			Member: model.MemberResponse{
				Model: member.Model,
				User: member.User,
				Team: member.Team,
				Name: member.Name,
				Role: member.Role,
			},
		})
	}
}