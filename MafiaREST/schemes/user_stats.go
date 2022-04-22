package schemes

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStats struct {
	UID          primitive.ObjectID `json:"uid" bson:"uid" binding:"required"`
	SessionCount uint64             `json:"session_count" bson:"session_count" binding:"required"`
	Wins         uint64             `json:"wins" bson:"wins" binding:"required"`
	Losses       uint64             `json:"losses" bson:"losses" binding:"required"`
	// supposed to be seconds
	TotalTime uint64 `json:"total_time" bson:"total_time" binding:"required"`
}
