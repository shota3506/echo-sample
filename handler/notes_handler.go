package handler

import (
	"../model"
	"github.com/labstack/echo"
	"net/http"
)

type noteParam struct {
	Title string
	Content string
	FolderId int `json:"folder_id"`
}

func (h *Handler) GetNote() echo.HandlerFunc {
	return func(c echo.Context) error {
		noteId := c.Param("id")
		note := model.Note{}
		result := h.DB.Preload("Folder").Preload("Member").First(&note, "id=?", noteId)
		if result.Error != nil { return h.return404(c, result.Error) }
		return c.JSON(http.StatusOK, struct {
			Note model.NoteResponse `json:"note"`
		} {
			Note: model.NoteResponse{
				Model: note.Model,
				Folder: note.Folder,
				Member: note.Member,
				Title: note.Title,
				Content: note.Content,
			},
		})
	}
}

func (h *Handler) CreateNote() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(noteParam)
		if err := c.Bind(param); err != nil { return h.return400(c, err) }
		folder := model.Folder{}
		result := h.DB.First(&folder, "id=?", param.FolderId)
		if result.Error != nil { return h.return404(c, result.Error) }
		e := h.setCurrentMember(c, folder.TeamID)
		if e != nil { return h.return404(c, e) }
		note := model.Note{
			Title: param.Title,
			Content: param.Content,
			Member: h.CurrentMember,
			Folder: folder,
		}

		if err := c.Validate(note); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":  err.Error(),
			})
		}
		h.DB.Create(&note)

		result = h.DB.Create(&note)
		if result.Error != nil { return h.return400(c, result.Error) }

		return c.JSON(http.StatusOK, struct {
			Note model.NoteResponse `json:"note"`
		} {
			Note: model.NoteResponse{
				Model: note.Model,
				Folder: note.Folder,
				Member: note.Member,
				Title: note.Title,
				Content: note.Content,
			},
		})
	}
}

func (h *Handler) UpdateNote() echo.HandlerFunc {
	return func(c echo.Context) error {
		noteId := c.Param("id")
		note := model.Note{}
		result := h.DB.Preload("Folder").Preload("Member").First(&note, "id=?", noteId)
		if result.Error != nil { return h.return404(c, result.Error) }

		param := new(noteParam)
		if err := c.Bind(param); err != nil { return h.return400(c, err) }

		note.Content = param.Content
		note.Title = param.Title
		if err := c.Validate(note); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":  err.Error(),
			})
		}
		h.DB.Save(&note)

		return c.JSON(http.StatusOK, struct {
			Note model.NoteResponse `json:"note"`
		} {
			Note: model.NoteResponse{
				Model: note.Model,
				Folder: note.Folder,
				Member: note.Member,
				Title: note.Title,
				Content: note.Content,
			},
		})
	}
}
