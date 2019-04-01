package model
import _ "gopkg.in/go-playground/validator.v9"

type Note struct {
	Model
	Folder Folder `json:"-"`
	FolderID uint `json:"folder_id"`
	Member Member `json:"-"`
	MemberID uint `json:"member_id"`
	Title string `gorm:"size:255" json:"title"`
	Content string `gorm:"size:255" json:"content"`
}

type NoteResponse struct {
	Model
	Folder Folder `json:"folder"`
	Member Member `json:"member"`
	Title string `json:"title"`
	Content string `json:"content"`
}
