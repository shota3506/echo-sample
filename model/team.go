package model

type Team struct {
	Model
	Name string `gorm:"unique_index" json:"name"`
	Members []Member `json:"members"`
}
