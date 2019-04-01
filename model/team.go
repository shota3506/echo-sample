package model

type Team struct {
	Model
	Name string `gorm:"unique_index" json:"name"`
	Members []Member `json:"members"`
	RootFolder Folder `json:"root_folder" gorm:"foreignkey:RootFolderID"`
	RootFolderID int `json:"root_folder_id"`
}
