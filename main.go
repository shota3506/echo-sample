package main

import (
	"./handler"
	"./model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(model.JWTTokenKey),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/" || c.Path() == "/login" || c.Path() == "/users" {
				return true
			}
			return true
		},
	}))

	source := os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@"+os.Getenv("DB_PROTOCOL")+"/"+os.Getenv("DB_NAME")+"?parseTime=true"
	db, err := gorm.Open("mysql", source)
	if err != nil {
		panic("データベースへの接続に失敗しました")
	}
	defer db.Close()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Member{})
	db.AutoMigrate(&model.Team{})

	h := &handler.Handler{DB: db}

	e.GET("/", h.Home())
	e.GET("/users/:id", h.GetUser())
	e.POST("/users", h.CreateUser())
	e.POST("/login", h.Login())

	e.GET("/teams", h.GetTeams())
	e.GET("/teams/:id", h.GetTeam())
	e.POST("/teams", h.CreateTeam())

	e.GET("/notes/:id", h.GetNote())
	e.POST("/notes", h.CreateNote())
	e.PUT("/notes/:id", h.UpdateNote())

	e.GET("/teams/:team_id/members", h.GetTeamMembers())
	e.POST("/teams/:team_id/members", h.CreateTeamMember())

	e.GET("/folders", h.GetFolders())
	e.GET("/folders/:id", h.GetFolder())
	e.PUT("/folders/:id", h.UpdateFolder())
	e.POST("/folders", h.CreateFolder())

	e.Logger.Fatal(e.Start(":1323"))
}
