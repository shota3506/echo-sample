package handler

import (
	"../model"
	"github.com/labstack/echo"
	"net/http"
)

type folderParam struct {
	Title string
	ParentId int `json:"parent_id"`
}

func (h *Handler) GetFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		folderId := c.Param("id")
		folder := model.Folder{}
		//result := h.DB.Preload("Folders", func(db *gorm.DB) *gorm.DB {
		//	return db.Where("tree_paths.length = ?", 1).Preload("Folders").Where("tree_paths.length = ?", 1)
		//}).First(&folder, "id=?", folderId)
		h.DB.Preload("Folders", "tree_paths.length = ?", 1).First(&folder, "id=?", folderId)
		result := h.DB.Preload("Folders.Folders", "tree_paths.length = ?", 1).First(&folder, "id=?", folderId)
		if result.Error != nil { return h.return404(c, result.Error) }
		return c.JSON(http.StatusOK, folder)
	}
}

func (h *Handler) GetFolders() echo.HandlerFunc {
	return func(c echo.Context) error {
		folders := []model.Folder{}
		result := h.DB.Preload("Folders","tree_paths.length = ?", 1).Preload("Folders.Folders", "tree_paths.length = ?", 1).Find(&folders)
		if result.Error != nil { return h.return404(c, result.Error) }
		return c.JSON(http.StatusOK, folders)
	}
}

func (h *Handler) CreateFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(folderParam)
		if err := c.Bind(param); err != nil { return h.return400(c, err) }
		folder := model.Folder{
			Title: param.Title,
		}
		if err := c.Validate(folder); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":  err.Error(),
			})
		}
		h.DB.Create(&folder)

		parent_tree_paths := []model.TreePath{}
		h.DB.Find(&parent_tree_paths, "descendant_id = ?", param.ParentId)

		for _, parent_tree_path := range parent_tree_paths {
			tree_path := model.TreePath{
				AncestorId: parent_tree_path.AncestorId,
				DescendantId: folder.ID,
				Length: parent_tree_path.Length + 1,
			}
			h.DB.Create(&tree_path)
		}

		tree_path := model.TreePath{
			AncestorId: folder.ID,
			DescendantId: folder.ID,
			Length: 0,
		}
		h.DB.Create(&tree_path)


		return c.JSON(http.StatusOK, folder)
	}
}

func (h *Handler) UpdateFolder() echo.HandlerFunc {
	return func(c echo.Context) error {
		folderId := c.Param("id")
		folder := model.Folder{}
		result := h.DB.First(&folder, "id=?", folderId)

		if result.Error != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"status": "Not Found",
			})
		}

		param := new(folderParam)
		if err := c.Bind(param); err != nil {
			return err
		}

		folder.Title = param.Title
		if err := c.Validate(folder); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":  err.Error(),
			})
		}
		h.DB.Save(&folder)

		return  c.JSON(http.StatusOK, folder)
	}
}
