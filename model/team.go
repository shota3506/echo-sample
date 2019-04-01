package model

type Team struct {
	Model
	Name string `gorm:"unique_index" json:"name" validate:"required"`
	Members []Member `json:"members"`
}
