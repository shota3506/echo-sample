package model

import "github.com/jinzhu/gorm"

type TreePath struct {
	gorm.Model
	AncestorId uint `json:"ancestor_id" validate:"required"`
	Ancestor Folder `json:"-"`
	DescendantId uint `json:"descendant_id" validate:"required"`
	Descendant Folder `json:"-"`
	Length int `json:"length" validate:"required"`
}

