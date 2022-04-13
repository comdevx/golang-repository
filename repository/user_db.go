package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryDB struct {
	collection *mongo.Collection
}

func NewUserRepositoryDB(db *mongo.Database) userRepositoryDB {
	collection := db.Collection("users")
	return userRepositoryDB{collection: collection}
}

func (r userRepositoryDB) GetAll() ([]User, error) {

	var result []User
	ctx := context.Background()
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r userRepositoryDB) Create(user User) (*User, error) {

	user.UserID = primitive.NewObjectID()
	ctx := context.Background()
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
