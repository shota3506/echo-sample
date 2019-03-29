package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	JWTTokenKey = "wkGRdkcF2taUE"
)

type User struct {
	Model
	Email string `gorm:"unique_index" json:"email"`
	Password string `json:"-"`
	WorkSpaces []WorkSpace `gorm:"many2many:user_work_spaces;" json:"work_spaces"`
}

func (u *User) IssueToken() (string, error)  {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = u.Email
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(JWTTokenKey))
}

