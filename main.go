package main

import (
	"project/handler"
	logs "project/helper"
	"project/middleware"
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

	public := router.Group("/api")
	{
		public.POST("/login", authenHandler.Login)
		public.POST("/change_password", authenHandler.ChangePassword)
		public.POST("/register", userHandler.NewUser)
	}

	user := router.Group("/api/users", middleware.Authorize)
	{
		user.GET("/", userHandler.GetUsers)
		user.GET("/:user_id", userHandler.GetUser)
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
