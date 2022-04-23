package main

import (
	"context"
	"project/handler"
	logs "project/helper"
	"project/repository"
	"project/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// router.POST("/users", userHandler.NewUser)

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

func initDatabase() *mongo.Database {

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://tongtest:tongtestgolang@cluster0.yc2a7.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		logs.Error(err)
		panic(err)
	}

	logs.Info("Database is connected")

	db := client.Database("golang")

	return db
}
