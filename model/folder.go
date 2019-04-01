package model
import _ "gopkg.in/go-playground/validator.v9"

type Folder struct {
	Model
	Title string `json:"title" validate:"required,min=1"`
	Team Team `json:"team"`
	TeamID uint `json:"team_id"`
	IsRoot bool `json:"is_root" gorm:"not null;default:false"`
	Folders []*Folder `gorm:"many2many:tree_paths;association_jointable_foreignkey:descendant_id;jointable_foreignkey:ancestor_id" json:"folders"`
}

//type Folder struct {
//	Model
//	Title string `json:"title"`
//	Folders []*Folder `gorm:"many2many:tree_paths;association_jointable_foreignkey:descendant_id" json:"folders"`
//}