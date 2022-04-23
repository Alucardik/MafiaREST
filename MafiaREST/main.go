package main

import (
	"MafiaREST/config"
	"MafiaREST/db"
	"MafiaREST/msgbroker"
	"MafiaREST/server"
	"MafiaREST/utils"
	"MafiaREST/worker"
	"flag"
	"log"
)

var (
	mode = flag.String("mode", config.SERVER_MODE, "server or worker mode")
)

func main() {
	dbHandle := db.CreateMongoDBHandle()
	err := dbHandle.InitConnection(config.DB_USERNAME, config.DB_PASS, config.DB_HOST, config.DB_PORT)
	utils.FailOnError("Failed to establish connection to the mongoDB", err)

	broker := msgbroker.CreateBroker()
	err = broker.InitConnection(config.QUEUE_HOST, config.QUEUE_PORT)
	utils.PanicOnError("Couldn't connect to RabbiMQ", err)
	defer broker.AbortConnection()

	q, err := broker.DeclareQueue(config.QUEUE_NAME)
	utils.PanicOnError("Couldn't declare work queue", err)

	flag.Parse()
	switch *mode {
	case config.SERVER_MODE:
		server.RunRESTServer(dbHandle, q)
	case config.WORKER_MODE:
		worker.Run(dbHandle, q)
	default:
		log.Fatalf("Unknown mode: %s; should be 'server' or 'worker'", *mode)
	}
}
