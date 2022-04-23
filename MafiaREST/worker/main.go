package worker

import (
	"MafiaREST/config"
	"MafiaREST/db"
	"MafiaREST/msgbroker"
	"MafiaREST/pdfgen"
	"MafiaREST/schemes"
	"MafiaREST/utils"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Run(handle db.MongoDbHandle, queue msgbroker.TaskQueue) {
	forever := make(chan bool)
	handler := func(task *msgbroker.Task) {
		pdfPath, err := pdfgen.GenReport(&task.User, &task.Stats)
		utils.NotifyOnError("Failed to produce pdf", err)
		fmt.Println("Saved report to:", pdfPath)
		uid := task.Stats.UID
		if err != nil {
			err = handle.UpdateReportMetaByUID(uid, &schemes.ReportMeta{
				UID:    uid,
				Status: schemes.REPORT_FAILED,
			})
		} else {
			file, err := os.Open(pdfPath)
			defer file.Close()
			defer os.Remove(pdfPath)

			utils.NotifyOnError("Couldn't open file", err)
			if err != nil {
				return
			}

			body := &bytes.Buffer{}
			_, err = io.Copy(body, file)
			utils.NotifyOnError("Couldn't open file", err)
			if err != nil {
				return
			}

			err = handle.UpdateReportMetaByUID(uid, &schemes.ReportMeta{
				UID:    uid,
				Status: schemes.REPORT_READY,
			})

			resp, err := http.Post(fmt.Sprintf("http://%s:%d%s/%s", config.REST_HOST, config.REST_PORT, config.UTILS_ENDPOINT, uid.String()[10:34]), "application/pdf", body)
			utils.NotifyOnError("Couldn't connect to the REST server", err)
			if err != nil || resp.StatusCode/100 != 2 {
				_ = handle.UpdateReportMetaByUID(uid, &schemes.ReportMeta{
					UID:    uid,
					Status: schemes.REPORT_FAILED,
				})
				return
			}
		}
		utils.NotifyOnError("Failed to update report meta", err)
	}

	err := queue.ConsumeTasks(&handler)
	utils.FailOnError("Failed to consume tasks", err)
	<-forever
}
