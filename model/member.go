package model
import _ "gopkg.in/go-playground/validator.v9"

type Member struct {
	Model
	User User `json:"user"`
	UserID int `json:"user_id" validate:"required"`
	Team Team `json:"team"`
	TeamID int `json:"team_id"  validate:"required"`
	Name string `json:"name" validate:"required,min=1"`
	Role string `json:"role" validate:"required"`
}

