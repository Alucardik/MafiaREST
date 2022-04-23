package schemes

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReportStatus uint8

const (
	REPORT_READY ReportStatus = iota
	REPORT_FAILED
	REPORT_NOT_FOUND
)

type ReportMeta struct {
	UID    primitive.ObjectID `bson:"uid"`
	Status ReportStatus       `bson:"status"`
	Path   string             `bson:"path"`
}
