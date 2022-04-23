package endpoints

import (
	"MafiaREST/config"
	"MafiaREST/msgbroker"
	"MafiaREST/schemes"
	"MafiaREST/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func (m *Manager) UpdateStats(ctx *gin.Context) {
	if !m.verify(ctx) {
		return
	}

	uid, err := primitive.ObjectIDFromHex(ctx.Param("uid"))
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fillInMsg(_INVALID_UID))
		return
	}

	var report schemes.SessionReport
	if err := ctx.BindJSON(&report); err != nil {
		ctx.JSON(http.StatusBadRequest, fillInMsg(_INVALID_REPORT))
		return
	}

	err = m.handle.UpdateUserStats(uid, &report)
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fillInMsg(_INVALID_UID))
		return
	}

	ctx.JSON(http.StatusOK, fillInMsg(_SUCCESS))
}

func (m *Manager) RequestStats(ctx *gin.Context) {
	if !m.verify(ctx) {
		return
	}

	uid, err := primitive.ObjectIDFromHex(ctx.Param("uid"))
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fillInMsg(_INVALID_UID))
		return
	}

	user, stats, err := m.handle.GetAggregatedStatsByUID(uid)
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fillInMsg(_GET_USER_ERR))
		return
	}

	err = m.taskQueue.PublishTask(msgbroker.Task{
		User:  *user,
		Stats: *stats,
	})
	utils.NotifyOnError("Failed to publish task", err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fillInMsg(_INTERNAL_ERR))
		return
	}

	ctx.JSON(http.StatusAccepted, fillInMsg(fmt.Sprintf("Your report will be available for download here: %s:%d%s/report/%s", config.REST_HOST, config.REST_PORT, config.STATS_ENDPOINT, uid.String()[10:34])))
}

func (m *Manager) GetStats(ctx *gin.Context) {
	if !m.verify(ctx) {
		return
	}

	uid, err := primitive.ObjectIDFromHex(ctx.Param("uid"))
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fillInMsg(_INVALID_UID))
		return
	}

	meta, err := m.handle.GetReportMetaByUID(uid)
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fillInMsg(_INTERNAL_ERR))
	}

	log.Println("IN GET REPORT", uid, meta, err)
	switch meta.Status {
	case schemes.REPORT_FAILED:
		ctx.JSON(http.StatusOK, fillInMsg(_FAILED_REPORT))
	case schemes.REPORT_NOT_FOUND:
		ctx.JSON(http.StatusNotFound, fillInMsg(_MISSING_REPORT))
	case schemes.REPORT_READY:
		log.Println("Sending pdf")
		ctx.FileAttachment(meta.Path, "MafiaPlayerReport.pdf")
	}
}
