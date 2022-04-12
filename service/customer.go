package service

import "go.mongodb.org/mongo-driver/bson/primitive"

type CustomerResponse struct {
	CustomerID primitive.ObjectID `bson:"_id" json:"id"`
	Name       string             `json:"name"`
	Status     bool               `json:"status"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(id string) (*CustomerResponse, error)
}
