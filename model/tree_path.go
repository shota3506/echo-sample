package model

import "github.com/jinzhu/gorm"

type TreePath struct {
	gorm.Model
	Ancestor Folder
	AncestorId uint `json:"ancestor_id"`
	DescendantId uint `json:"descendant_id"`
	Descendant Folder
	Length int `json:"length"`
}

