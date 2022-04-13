package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type customerRepositoryDB struct {
	collection *mongo.Collection
}

func NewCustomerRepositoryDB(db *mongo.Database) customerRepositoryDB {
	collection := db.Collection("customers")
	return customerRepositoryDB{
		collection: collection,
	}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {

	var result []Customer
	ctx := context.Background()
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r customerRepositoryDB) GetByID(id string) (*Customer, error) {

	var result Customer
	ctx := context.Background()
	ConvID, _ := primitive.ObjectIDFromHex(id)
	err := r.collection.FindOne(ctx, bson.M{"_id": ConvID}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
