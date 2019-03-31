package model

type Member struct {
	Model
	User User `json:"-"`
	UserID int `json:"user_id"`
	Team Team `json:"-"`
	TeamID int `json:"team_id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

