package endpoints

import (
	"MafiaREST/db"
	"MafiaREST/msgbroker"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Manager struct {
	isSet     bool
	handle    db.MongoDbHandle
	taskQueue msgbroker.TaskQueue
}

func (m *Manager) verify(ctx *gin.Context) bool {
	if !m.isSet {
		ctx.JSON(http.StatusInternalServerError, fillInMsg(_INTERNAL_ERR))
	}

	return m.isSet
}

func (m *Manager) Setup(dbHandle db.MongoDbHandle, queue msgbroker.TaskQueue) {
	m.handle = dbHandle
	m.taskQueue = queue
	m.isSet = true
}
