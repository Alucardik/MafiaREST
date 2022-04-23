package server

import (
	"MafiaREST/config"
	"MafiaREST/db"
	"MafiaREST/endpoints"
	"MafiaREST/msgbroker"
	"MafiaREST/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RunRESTServer(handle db.MongoDbHandle, queue msgbroker.TaskQueue) {
	manager := endpoints.Manager{}
	manager.Setup(handle, queue)

	router := gin.Default()
	// TODO: add right-controls via gin.Group (see example on using middleware)
	router.POST(config.USERS_ENDPOINT, manager.AddUser)
	router.GET(config.USERS_ENDPOINT+"/:id", manager.GetUser)
	router.GET(config.USERS_ENDPOINT, manager.GetUsers)
	router.PATCH(config.USERS_ENDPOINT+"/:id", manager.UpdateUser)
	router.DELETE(config.USERS_ENDPOINT+"/:id", manager.DeleteUser)

	router.PUT(config.STATS_ENDPOINT+"/:uid", manager.UpdateStats)
	router.GET(config.STATS_ENDPOINT+"/report/:uid", manager.GetStats)
	router.GET(config.STATS_ENDPOINT+"/:uid", manager.RequestStats)

	router.POST(config.UTILS_ENDPOINT+"/:uid", manager.SaveReport)

	err := router.Run(fmt.Sprintf("%s:%d", config.REST_HOST, config.REST_PORT))
	utils.PanicOnError("", err)
}
