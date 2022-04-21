package endpoints

import (
	"MafiaREST/schemes"
	"MafiaREST/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func UpdateStats(ctx *gin.Context) {
	if !verifyHandle(ctx) {
		return
	}

	uid, err := primitive.ObjectIDFromHex(ctx.Param("uid"))
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fillInError(_INVALID_UID))
		return
	}

	var report schemes.SessionReport
	if err := ctx.BindJSON(&report); err != nil {
		ctx.JSON(http.StatusBadRequest, fillInError(_INVALID_REPORT))
		return
	}

	fmt.Println(report)

	err = handle.UpdateUserStats(uid, &report)
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fillInError(_INVALID_UID))
		return
	}

	ctx.JSON(http.StatusOK, fillInError(_SUCCESS))
}
