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

func (dh *mongoDbHandle) GetStatsByUID(uid primitive.ObjectID) (*schemes.UserStats, error) {
	res := schemes.UserStats{}
	ctx, cancel := context.WithTimeout(context.Background(), _REQUEST_TM)
	defer cancel()

	err := dh.stats.FindOne(ctx, bson.D{{"uid", uid}}).Decode(&res)
	return &res, err
}

func (dh *mongoDbHandle) GetAggregatedStatsByUID(uid primitive.ObjectID) (*schemes.User, *schemes.UserStats, error) {
	user, err := dh.GetUserById(uid)
	if err != nil {
		return nil, nil, err
	}

	stats, err := dh.GetStatsByUID(uid)
	if err != nil {
		return nil, nil, err
	}

	return user, stats, nil
}
