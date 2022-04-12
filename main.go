package main

import (
	"bank/handler"
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

	customerRepository := repository.NewCustomerRepositoryDB((*mongo.Database)(db))
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	_ = customerRepository

	customerService := service.NewCustomerService(customerRepositoryMock)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := gin.Default()

	router.GET("/customers", customerHandler.GetCustomers)
	router.GET("/customers/:customer_id", customerHandler.GetCustomer)

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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://tongtest:tongtestgolang@cluster0.yc2a7.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	db := client.Database("golang")

	return db
}
