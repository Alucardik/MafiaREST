package server

import (
	"MafiaREST/config"
	"MafiaREST/db"
	"MafiaREST/endpoints"
	"MafiaREST/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RunRESTServer(handle db.MongoDbHandle) {
	endpoints.SetupDbHandle(handle)

	router := gin.Default()
	// TODO: add right-controls via gin.Group (see example on using middleware)
	router.POST(config.USERS_ENDPOINT, endpoints.AddUser)
	router.GET(config.USERS_ENDPOINT+"/:id", endpoints.GetUser)
	router.GET(config.USERS_ENDPOINT, endpoints.GetUsers)
	router.PATCH(config.USERS_ENDPOINT+"/:id", endpoints.UpdateUser)
	router.DELETE(config.USERS_ENDPOINT+"/:id", endpoints.DeleteUser)

	router.PUT(config.STATS_ENDPOINT+"/:uid", endpoints.UpdateStats)
	// TODO: publish task for the worker
	router.GET(config.STATS_ENDPOINT+"/:uid", endpoints.)
	// TODO: donwload pdf afer worker is finished
	router.GET(config.STATS_ENDPOINT + "/report/:uid", endpoints.)

	err := router.Run(fmt.Sprintf("%s:%d", config.REST_HOST, config.REST_PORT))
	utils.PanicOnError("", err)
}
