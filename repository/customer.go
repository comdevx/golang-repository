package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	CustomerID primitive.ObjectID `bson:"_id" json:"id"`
	Name       string             `json:"name"`
	DOB        time.Time          `json:"dob"`
	City       string             `json:"city"`
	ZipCode    string             `json:"zipcode"`
	Status     bool               `json:"status"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetByID(id string) (*Customer, error)
}
