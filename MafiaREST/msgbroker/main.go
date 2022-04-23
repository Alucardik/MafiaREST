package msgbroker

type MessageBroker interface {
	InitConnection(address string, port int) error
	AbortConnection()
	DeclareQueue(name string) (TaskQueue, error)
}

func CreateBroker() MessageBroker {
	return &rabbitMQBroker{}
}
