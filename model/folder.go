package model


type Folder struct {
	Model
	Title string `json:"title"`
	Team Team `json:"team"`
	TeamID uint `json:"team_id"`
	Folders []*Folder `gorm:"many2many:tree_paths;association_jointable_foreignkey:descendant_id;jointable_foreignkey:ancestor_id" json:"folders"`
}

//type Folder struct {
//	Model
//	Title string `json:"title"`
//	Folders []*Folder `gorm:"many2many:tree_paths;association_jointable_foreignkey:descendant_id" json:"folders"`
//}