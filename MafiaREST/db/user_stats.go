package db

import (
	"MafiaREST/schemes"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dh *mongoDbHandle) UpdateUserStats(uid primitive.ObjectID, update *schemes.SessionReport) error {
	ctx, cancel := context.WithTimeout(context.Background(), _REQUEST_TM)
	defer cancel()
	updateInfo := bson.D{}

	switch update.Outcome {
	case schemes.SESSION_LOSE:
		updateInfo = bson.D{{
			"$inc",
			bson.D{
				{"session_count", 1},
				{"losses", 1},
				{"total_time", update.Duration},
			}}}
	case schemes.SESSION_WIN:
		updateInfo = bson.D{{
			"$inc",
			bson.D{
				{"session_count", 1},
				{"wins", 1},
				{"total_time", update.Duration},
			}}}

	default:
		return errors.New("unknown outcome")
	}

	_, err := dh.stats.UpdateOne(ctx, bson.D{{"uid", uid}}, updateInfo)
	return err
}

func (dh *mongoDbHandle) DeleteUserStatsByUID(uid primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), _REQUEST_TM)
	defer cancel()

	_, err := dh.stats.DeleteOne(ctx, bson.D{{"uid", uid}})
	return err
}