package model

type Note struct {
	Model
	Folder Folder `json:"folder"`
	FolderID uint `json:"folder_id"`
	Content string `gorm:"size:255" json:"content"`
	Title string `gorm:"size:255" json:"title"`
}
