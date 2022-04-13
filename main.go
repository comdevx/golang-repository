package main

import (
	"bank/handler"
	"bank/logs"
	"bank/repository"
	"bank/service"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	initTimeZone()
	db := initDatabase()

	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	_ = customerRepositoryMock

	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := gin.Default()

	router.GET("/customers", customerHandler.GetCustomers)
	router.GET("/customers/:customer_id", customerHandler.GetCustomer)

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

	logs.Info("Mongo is connected")

	db := client.Database("golang")

	return db
}
