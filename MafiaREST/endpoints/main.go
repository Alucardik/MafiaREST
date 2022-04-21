package endpoints

import (
	"MafiaREST/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

var handle db.MongoDbHandle = nil

func verifyHandle(ctx *gin.Context) bool {
	if handle == nil {
		ctx.JSON(http.StatusInternalServerError, fillInError(_INTERNAL_ERR))
	}

	return handle != nil
}

func SetupDbHandle(dbHandle db.MongoDbHandle) {
	handle = dbHandle
}
