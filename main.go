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

	db, err := gorm.Open("mysql", os.Getenv("DATABASE_SOURCE"))
	if err != nil {
		panic("データベースへの接続に失敗しました")
	}
	defer db.Close()
	db.AutoMigrate(&model.User{})

	h := &handler.Handler{DB: db}

	r := e.Group("")
	r.Use(middleware.JWT([]byte("wkGRdkcF2taUE")))

	e.GET("/", h.Home())
	r.GET("/users/:id", h.GetUsers())
	e.POST("/users", h.SaveUser())
	e.POST("/login", h.Login())

	e.Logger.Fatal(e.Start(":1323"))
}
