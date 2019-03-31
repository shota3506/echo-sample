package model

type Note struct {
	Model
	Content string `gorm:"size:255" json:"content"`
	Title string `gorm:"size:255" json:"title"`
}
