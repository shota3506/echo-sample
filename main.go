package main

import (
	"net/http"
	"os"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"./handler"
	"./model"
)

func main() {
	e := echo.New()

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
			return false
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

	e.GET("/teams/:id", h.GetTeam())
	e.POST("/teams", h.CreateTeam())

	e.GET("/notes/:id", h.GetNote())
	e.POST("/notes", h.CreateNote())
	e.PUT("/notes/:id", h.UpdateNote())

	e.Logger.Fatal(e.Start(":1323"))
}
