package msgbroker

type MessageBroker interface {
	InitConnection(address string, port int)
	AbortConnection()
	PublishTask()
}