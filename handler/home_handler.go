package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) Home() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"Greet": "World World!",
		})
	}
}