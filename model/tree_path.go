package model

import "github.com/jinzhu/gorm"

type TreePath struct {
	gorm.Model
	Ancestor Folder
	AncestorId uint `json:"ancestor_id"`
	Descendant Folder
	DescendantId uint `json:"descendant_id"`
	Length int `json:"length"`
}

