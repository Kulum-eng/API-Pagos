package ports

type Broker interface {
	Connect() error
	InitChannel(queueueName string) error
	Publish(message string) error
}