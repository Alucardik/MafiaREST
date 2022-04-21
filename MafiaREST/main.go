package main

import (
	"MafiaREST/config"
	"MafiaREST/db"
	"MafiaREST/server"
	"MafiaREST/utils"
)

func main() {
	dbHandle := db.CreateMongoDBHandle()
	err := dbHandle.InitConnection(config.DB_USERNAME, config.DB_PASS, config.DB_HOST, config.DB_PORT)
	utils.FailOnError("Failed to establish connection to the mongoDB", err)

	server.RunRESTServer(dbHandle)
}
