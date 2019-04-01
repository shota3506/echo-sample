package model
import _ "gopkg.in/go-playground/validator.v9"

type Note struct {
	Model
  Folder Folder `json:"folder"`
	FolderID uint `json:"folder_id validate:"required""`
	Content string `gorm:"size:255" json:"content" validate:"required"`
	Title string `gorm:"size:255" json:"title" validate:"required"`
}
