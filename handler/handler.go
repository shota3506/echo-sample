package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
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