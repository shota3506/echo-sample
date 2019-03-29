package model

type Team struct {
	Model
	Name string `gorm:"unique_index" json:"name"`
	Users []User `gorm:"many2many:user_teams;" json:"users"`
}
