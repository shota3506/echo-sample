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

func (h *Handler) getCurrentUser(c echo.Context) (model.User, error) {
	userEmail := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["email"].(string)
	currentUser := model.User{}
	result := h.DB.Preload("Members").Preload("Teams").First(&currentUser, "email=?", userEmail)
	return currentUser, result.Error
}

func (h *Handler) getCurrentMember(c echo.Context, teamID uint) (model.Member, error) {
	currentUser, e := h.getCurrentUser(c)
	if e != nil { return model.Member{}, e }
	currentMember := model.Member{}
	result := h.DB.Preload("User").Preload("Team").First(&currentMember, "user_id=? AND team_id=?", currentUser.ID, teamID)
	return currentMember, result.Error
}
