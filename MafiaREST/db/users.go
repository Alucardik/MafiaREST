package db

import (
	"MafiaREST/schemes"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (dh *mongoDbHandle) AddUser(user *schemes.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), _REQUEST_TM)
	defer cancel()

	res, err := dh.users.InsertOne(ctx, user)
	return res, err
}

func (dh *mongoDbHandle) GetUserById(id primitive.ObjectID) (*schemes.User, error) {
	res := schemes.User{}
	ctx, cancel := context.WithTimeout(context.Background(), _REQUEST_TM)
	defer cancel()

	err := dh.users.FindOne(ctx, bson.D{{"_id", id}}).Decode(&res)
	return &res, err
}

func (dh *mongoDbHandle) GetAllUsers() (*[]schemes.User, error) {
	var res []schemes.User
	ctx, cancel := context.WithTimeout(context.Background(), _REQUEST_TM)
	defer cancel()

	iterator, err := dh.users.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	// TODO: is it ok to use the same ctx here??
	if err = iterator.All(ctx, &res); err != nil {
		return nil, err
	}

	return &res, err
}

func (dh *mongoDbHandle) UpdateUserById(id primitive.ObjectID, updated *schemes.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), _REQUEST_TM)
	defer cancel()

	_, err := dh.users.ReplaceOne(ctx, bson.D{{"_id", id}}, updated)
	return err
}

func (dh *mongoDbHandle) DeleteUserById(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), _REQUEST_TM)
	defer cancel()

	_, err := dh.users.DeleteOne(ctx, bson.D{{"_id", id}})
	return err
}
