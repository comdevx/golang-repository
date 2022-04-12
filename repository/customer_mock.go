package repository

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{CustomerID: primitive.NewObjectID(), Name: "test1", City: "city1", ZipCode: "zipcode1", DOB: time.Now(), Status: false},
		{CustomerID: primitive.NewObjectID(), Name: "test2", City: "city2", ZipCode: "zipcode2", DOB: time.Now(), Status: false},
		{CustomerID: primitive.NewObjectID(), Name: "test3", City: "city3", ZipCode: "zipcode3", DOB: time.Now(), Status: false},
	}

	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetByID(id string) (*Customer, error) {

	convID, _ := primitive.ObjectIDFromHex(id)
	for _, customer := range r.customers {
		if customer.CustomerID == convID {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}
