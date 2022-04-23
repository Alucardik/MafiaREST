package db

import (
	"MafiaREST/schemes"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
)

func (dh *mongoDbHandle) UpdateReportMetaByUID(uid primitive.ObjectID, updated *schemes.ReportMeta) error {
	ctx, cancel := context.WithTimeout(context.Background(), _REQUEST_TM)
	defer cancel()

	options := options2.Replace().SetUpsert(true)
	_, err := dh.reports.ReplaceOne(ctx, bson.D{{"uid", uid}}, updated, options)
	return err
}

func (dh *mongoDbHandle) GetReportMetaByUID(uid primitive.ObjectID) (*schemes.ReportMeta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), _REQUEST_TM)
	defer cancel()

	report := schemes.ReportMeta{Status: schemes.REPORT_NOT_FOUND}
	err := dh.reports.FindOne(ctx, bson.D{{"uid", uid}}).Decode(&report)
	if err == mongo.ErrNoDocuments {
		err = nil
	}

	return &report, err
}
