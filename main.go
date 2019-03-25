package main

import (
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
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/" || c.Path() == "/login" || c.Path() == "/users" {
				return true
			}
			return false
		},
	}))

	db, err := gorm.Open("mysql", os.Getenv("DATABASE_SOURCE"))
	if err != nil {
		panic("データベースへの接続に失敗しました")
	}
	defer db.Close()
	db.AutoMigrate(&model.User{})

	h := &handler.Handler{DB: db}

	e.GET("/", h.Home())
	e.GET("/users/:id", h.GetUser())
	e.POST("/users", h.SaveUser())
	e.POST("/login", h.Login())

	e.Logger.Fatal(e.Start(":1323"))
}
