package db

import (
	"MafiaREST/schemes"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbHandle interface {
	InitConnection(username, password, host string, port int) error

	AddUser(user *schemes.User) (*mongo.InsertOneResult, error)
	GetUserById(id primitive.ObjectID) (*schemes.User, error)
	GetAllUsers() (*[]schemes.User, error)
	UpdateUserById(id primitive.ObjectID, updated *schemes.User) error
	DeleteUserById(id primitive.ObjectID) error

	UpdateUserStats(uid primitive.ObjectID, update *schemes.SessionReport) error
	DeleteUserStatsByUID(uid primitive.ObjectID) error
	GetAggregatedStatsByUID(uid primitive.ObjectID) (*schemes.User, *schemes.UserStats, error)

	UpdateReportMetaByUID(uid primitive.ObjectID, updated *schemes.ReportMeta) error
	GetReportMetaByUID(uid primitive.ObjectID) (*schemes.ReportMeta, error)
}

type mongoDbHandle struct {
	ctx     context.Context
	client  *mongo.Client
	mafiaDb *mongo.Database
	users   *mongo.Collection
	stats   *mongo.Collection
	reports *mongo.Collection
}

func (dh *mongoDbHandle) InitConnection(username, password, host string, port int) error {
	ctx, cancel := context.WithTimeout(context.Background(), _CONNECTION_TM)
	defer cancel()

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d/", username, password, host, port))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	dh.client = client

	err = dh.client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	dh.mafiaDb = dh.client.Database("mafia")
	dh.users = dh.mafiaDb.Collection("users")
	_, err = dh.users.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{"email", 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}

	dh.stats = dh.mafiaDb.Collection("stats")
	dh.reports = dh.mafiaDb.Collection("reports")
	return nil
}

func CreateMongoDBHandle() MongoDbHandle {
	return &mongoDbHandle{}
}
