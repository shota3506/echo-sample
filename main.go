package main

import (
	"./handler"
	"./model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db, err := gorm.Open("mysql", os.Getenv("DATABASE_SOURCE"))
	if err != nil {
		panic("データベースへの接続に失敗しました")
	}
	defer db.Close()
	db.AutoMigrate(&model.User{})

	r := e.Group("")
	r.Use(middleware.JWT([]byte("wkGRdkcF2taUE")))

	e.GET("/", handler.Home())
	r.GET("/users/:id", handler.GetUsers())
	e.POST("/users", handler.SaveUser())
	e.POST("/login", handler.Login())

	e.Logger.Fatal(e.Start(":1323"))
}
