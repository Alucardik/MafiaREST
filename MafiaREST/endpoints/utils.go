package endpoints

import (
	"MafiaREST/config"
	"MafiaREST/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
)

func (m *Manager) SaveReport(ctx *gin.Context) {
	if !m.verify(ctx) {
		return
	}

	uid, err := primitive.ObjectIDFromHex(ctx.Param("uid"))
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fillInMsg(_INVALID_UID))
		return
	}

	buf, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "")
		return
	}

	filePath := config.TMP_FILE_PATH + "/" + uid.String()[10:34] + "_report.pdf"

	err = os.WriteFile(filePath, buf, 0666)
	utils.NotifyOnError("Failed to copy report", err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	ctx.JSON(http.StatusOK, "")
}
