package model

type User struct {
	Model
	Name string `gorm:"unique_index" json:"name"`
	Password string `json:"-"`
	WorkSpaces []WorkSpace `gorm:"many2many:user_work_spaces;" json:"work_spaces"`
}