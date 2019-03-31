package model

type Member struct {
	Model
	User User `json:"user"`
	UserID int
	Team Team `json:"team"`
	TeamID int
	Name string `json:"name"`
}

