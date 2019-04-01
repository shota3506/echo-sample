package model
import _ "gopkg.in/go-playground/validator.v9"

type Note struct {
	Model
	Content string `gorm:"size:255" json:"content" validate:"required"`
	Title string `gorm:"size:255" json:"title" validate:"required"`
}
