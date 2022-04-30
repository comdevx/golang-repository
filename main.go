package main

import (
	"project/handler"
	logs "project/helper"
	"project/repository"
	"project/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	initData()
	initTimeZone()
	db := initDatabase()

	userRepository := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	authenService := service.NewAuthenService(userRepository)
	authenHandler := handler.NewAuthenHandler(authenService)

	router := gin.Default()

	authen := router.Group("/api/authen")
	{
		authen.POST("/login", authenHandler.Login)
	}

	user := router.Group("/api/users")
	{
		user.GET("/", userHandler.GetUsers)
		user.GET("/:user_id", userHandler.GetUser)
		user.POST("/", userHandler.NewUser)
	}

	logs.Info("Started port 3000")
	router.Run(":3000")
}

func initTimeZone() {

	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		logs.Error(err)
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *gorm.DB {

	// dsn := "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(sqlite.Open("./db/test.db"), &gorm.Config{})
	if err != nil {
		logs.Error(err)
		panic(err)
	}

	logs.Info("Database is connected")

	return db
}

func initData() {

	err := godotenv.Load()
	if err != nil {
		logs.Error(err)
		panic("Error loading .env file")
	}
}
