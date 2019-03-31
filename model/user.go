package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)


type User struct {
	Model
	Email string `gorm:"unique_index" json:"email"`
	Password string `json:"-"`
	Members []Member `json:"members"`
}

const (
	JWTTokenKey = "wkGRdkcF2taUE"
)

func (u *User) IssueToken() (string, error)  {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = u.Email
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(JWTTokenKey))
}

