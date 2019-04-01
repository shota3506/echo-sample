package handler

import (
	"../model"
	"github.com/labstack/echo"
	"net/http"
)

type teamParam struct {
	Name string `json:"name"`
	MemberName string `json:"member_name"`
}

func (h *Handler) GetTeams() echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser, e := h.getCurrentUser(c)
		if e != nil { return h.return404(c, e) }
		return c.JSON(http.StatusOK, struct {
			Teams []model.Team `json:"teams"`
		} {
			Teams: currentUser.Teams,
		})
	}
}

func (h *Handler) GetTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("id")
		team := model.Team{}
		result := h.DB.Preload("Users").First(&team, "id=?", teamId)
		if result.Error != nil { return h.return404(c, result.Error) }
		return c.JSON(http.StatusOK, struct {
			Team model.Team `json:"team"`
		} {
			Team: team,
		})
	}
}

func (h *Handler) CreateTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		currentUser, e := h.getCurrentUser(c)
		if e != nil { return h.return404(c, e) }

		param := new(teamParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		team := model.Team{
			Name: param.Name,
		}
    
		if err := c.Validate(team); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":  err.Error(),
			})
		}
		
		result := h.DB.Create(&team)
		
    if result.Error != nil { return h.return400(c, result.Error) }
		folder := model.Folder{
			Title: "root",
			IsRoot: true,
			Team: team,
		}
		result = h.DB.Create(&folder)

		member := model.Member{
			User: currentUser,
			Team: team,
			Name: param.MemberName,
			Role: "admin",
		}
		if err := c.Validate(member); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":  err.Error(),
			})
		}
		result = h.DB.Create(&member)
		if result.Error != nil { return h.return400(c, result.Error) }
		return c.JSON(http.StatusOK, struct {
			Team model.Team `json:"team"`
		} {
			Team: team,
		})
	}
}
