package worker

import (
	"MafiaREST/db"
	"MafiaREST/msgbroker"
	"MafiaREST/pdfgen"
	"MafiaREST/schemes"
	"MafiaREST/utils"
	"fmt"
)

func Run(handle db.MongoDbHandle, queue msgbroker.TaskQueue) {
	forever := make(chan bool)
	handler := func(task *msgbroker.Task) {
		pdfPath, err := pdfgen.GenReport(&task.User, &task.Stats)
		utils.NotifyOnError("Failed to produce pdf", err)
		fmt.Println("IN WORKER", pdfPath)
		uid := task.Stats.UID
		if err != nil {
			err = handle.UpdateReportMetaByUID(uid, &schemes.ReportMeta{
				UID:    uid,
				Status: schemes.REPORT_FAILED,
				Path:   "",
			})
		} else {
			err = handle.UpdateReportMetaByUID(uid, &schemes.ReportMeta{
				UID:    uid,
				Status: schemes.REPORT_READY,
				Path:   pdfPath,
			})
		}
		utils.NotifyOnError("Failed to update report meta", err)
	}

	err := queue.ConsumeTasks(&handler)
	utils.FailOnError("Failed to consume tasks", err)
	<-forever
}
