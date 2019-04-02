package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"../model"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) return400(c echo.Context, e error) error {
	return c.JSON(http.StatusBadRequest, map[string]error{
		"error": e,
	})
}

func (h *Handler) return404(c echo.Context, e error) error {
	return c.JSON(http.StatusNotFound, map[string]error{
		"error": e,
	})
}

func (h *Handler) setCurrentUser(u *model.User, c echo.Context) error {
	userEmail := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
	result := h.DB.Preload("Members").Preload("Teams").First(&u, "email=?", userEmail)
	return result.Error
}

func (h *Handler) setCurrentMember(m *model.Member, c echo.Context, teamID uint) error {
	currentUser := model.User{}
	e := h.setCurrentUser(&currentUser, c)
	if e != nil { return e }
	result := h.DB.Preload("User").Preload("Team").First(&m, "user_id=? AND team_id=?", currentUser.ID, teamID)
	return result.Error
}
