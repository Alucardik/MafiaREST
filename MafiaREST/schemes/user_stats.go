package schemes

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStats struct {
	UID          primitive.ObjectID `json:"uid" bson:"uid"`
	SessionCount uint64             `json:"session_count" bson:"session_count"`
	WinRate      uint64             `json:"win_rate" bson:"win_rate"`
	LossRate     uint64             `json:"loss_rate" bson:"loss_rate"`
	TotalTime    uint64             `json:"total_time" bson:"total_time"`
}
