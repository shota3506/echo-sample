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
		e := h.setCurrentUser(c)
		if e != nil { return h.return404(c, e) }
		return c.JSON(http.StatusOK, struct {
			Teams []model.Team `json:"teams"`
		} {
			Teams: h.CurrentUser.Teams,
		})
	}
}

func (h *Handler) GetTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		teamId := c.Param("id")
		team := model.Team{}
		result := h.DB.Preload("Members").First(&team, "id=?", teamId)
		if result.Error != nil { return h.return404(c, result.Error) }
		return c.JSON(http.StatusOK, struct {
			Team model.TeamResponse `json:"team"`
		} {
			Team: model.TeamResponse{
				Model: team.Model,
				Name: team.Name,
				Members: team.Members,
				RootFolder: team.GetRootFolder(h.DB),
			},
		})
	}
}

func (h *Handler) CreateTeam() echo.HandlerFunc {
	return func(c echo.Context) error {
		e := h.setCurrentUser(c)
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

		tree_path := model.TreePath{
			AncestorId: folder.ID,
			DescendantId: folder.ID,
			Length: 0,
		}
		h.DB.Create(&tree_path)

		member := model.Member{
			User: h.CurrentUser,
			Name: param.MemberName,
			Role: "admin",
			Team: team,
		}

		if err := c.Validate(member); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":  err.Error(),
			})
		}
		h.DB.Model(&team).Association("Members").Append(member)
		return c.JSON(http.StatusOK, struct {
			Team model.TeamResponse `json:"team"`
		} {
			Team: model.TeamResponse{
				Model: team.Model,
				Name: team.Name,
				Members: team.Members,
				RootFolder: folder,
			},
		})
	}
}
