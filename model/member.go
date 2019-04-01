package model
import _ "gopkg.in/go-playground/validator.v9"

type Member struct {
	Model
	User User `json:"-"`
	UserID uint `json:"user_id"`
	Team Team `json:"-"`
	TeamID uint `json:"team_id"`
	Name string `json:"name" validate:"required,min=1"`
	Role string `json:"role" validate:"required"`
}

type MemberResponse struct {
	Model
	User User `json:"user"`
	Team Team `json:"team"`
	Name string `json:"name"`
	Role string `json:"role"`
}

