package model

type Team struct {
	Model
	Name string `gorm:"unique_index" json:"name" validate:"required"`
	Members []Member `json:"members"`
	Folders []Folder `json:"folders"`
}

//func (t *Team) GetRootFolder(db *gorm.DB) {
//
//}