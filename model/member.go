package model

type Member struct {
	Model
	User User `json:"user"`
	UserID int `json:"user_id"`
	Team Team `json:"team"`
	TeamID int `json:"team_id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

