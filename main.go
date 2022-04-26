package main

import (
	"project/handler"
	logs "project/helper"
	"project/repository"
	"project/service"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	initTimeZone()
	db := initDatabase()

	userRepository := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	router.GET("/users", userHandler.GetUsers)
	router.GET("/users/:user_id", userHandler.GetUser)
	router.POST("/users", userHandler.NewUser)

	logs.Info("Started port 3000")
	router.Run(":3000")
}

func initTimeZone() {

	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *gorm.DB {

	// dsn := "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		logs.Error(err)
		panic(err)
	}

	logs.Info("Database is connected")

	return db
}
