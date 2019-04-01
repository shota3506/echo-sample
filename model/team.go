package model

import "github.com/jinzhu/gorm"

type Team struct {
	Model
	Name string `gorm:"unique_index" json:"name" validate:"required"`
	Members []Member `json:"-"`
	Folders []Folder `json:"-"`
}

type TeamResponse struct {
	Model
	Name string `json:"name"`
	Members []Member `json:"members"`
	RootFolder Folder `json:"root_folder"`
}

func (t *Team) GetRootFolder(db *gorm.DB) Folder {
	rootFolder := Folder{}
	db.First(&rootFolder, "team_id=? and is_root=?", t.ID, true)
	return rootFolder
}