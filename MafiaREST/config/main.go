package config

const (
	SERVER_MODE = "server"
	WORKER_MODE = "worker"
)

const (
	USERS_ENDPOINT = "/users"
	STATS_ENDPOINT = "/stats"
	UTILS_ENDPOINT = "/utils"
)

// TODO: add getenv as well
const (
	DB_PORT    = 27017
	QUEUE_PORT = 5672
	REST_PORT  = 8080
)

var (
	DB_HOST     = setEnvConfigParam("DB_HOST", "localhost", true)
	DB_USERNAME = setEnvConfigParam("DB_USERNAME", "root", true)
	DB_PASS     = setEnvConfigParam("DB_PASS", "example", true)
)

var (
	QUEUE_HOST = setEnvConfigParam("QUEUE_HOST", "localhost", true)
	QUEUE_NAME = setEnvConfigParam("QUEUE_NAME", "worker_queue", true)
)

var (
	REST_HOST = setEnvConfigParam("REST_HOST", "localhost", true)
)

var (
	TMP_FILE_PATH = setEnvConfigParam("TMP_FILE_PATH", "/tmp", true)
)
