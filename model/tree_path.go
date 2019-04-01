package model

import "github.com/jinzhu/gorm"

type TreePath struct {
	gorm.Model
	Ancestor Folder
	AncestorId uint `json:"ancestor_id" validate:"required"`
	DescendantId uint `json:"descendant_id" validate:"required"`
	Descendant Folder
	Length int `json:"length" validate:"required"`
}

