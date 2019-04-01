package handler

import (
	"../model"
	"github.com/labstack/echo"
	"net/http"
)

type folderParam struct {
	Title string
}

func (h *Handler) GetFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		folderId := c.Param("id")
		folder := model.Folder{}
		result := h.DB.Where("tree_paths.length = ?", 1).Preload("Folders").First(&folder, "id=?", folderId)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"status": "Not Found",
			})
		}
		return c.JSON(http.StatusOK, folder)
	}
}

func (h *Handler) GetFolders() echo.HandlerFunc {
	return func(c echo.Context) error {
		folders := []model.Folder{}
		result := h.DB.Preload("Folders","tree_paths.length = ?", 1).Preload("Folders").Find(&folders)

		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"status": "Not Found",
			})
		}
		return c.JSON(http.StatusOK, folders)
	}
}

func (h *Handler) UpdateFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		folders := []model.Folder{}
		result := h.DB.Preload("Folders","tree_paths.length = ?", 1).Preload("Folders").Find(&folders)

		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"status": "Not Found",
			})
		}
		return c.JSON(http.StatusOK, folders)
	}
}