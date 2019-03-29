package model

type WorkSpace struct {
	Model
	Name string `gorm:"unique_index" json:"name"`
	Users []User `gorm:"many2many:user_work_spaces;" json:"users"`
}
