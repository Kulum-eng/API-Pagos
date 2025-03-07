package adapters

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"

)

type RabbitMQBroker struct {
	Host        string
	Port        int
	User        string
	Password    string
	Conn        *amqp091.Connection
	Channel     *amqp091.Channel
	QueueueName string
}

func NewRabbitMQBroker(host string, port int, user, password string) *RabbitMQBroker {
	return &RabbitMQBroker{Host: host, Port: port, User: user, Password: password}
}

func (b *RabbitMQBroker) Connect() error {
	var err error
	stringConnection := fmt.Sprintf("amqp://%s:%s@%s:%d/", b.User, b.Password, b.Host, b.Port)

	b.Conn, err = amqp091.Dial(stringConnection)
	if err != nil {
		return err
	}

	fmt.Println("✅ Conectado a RabbitMQ")

	return nil
}

func (b *RabbitMQBroker) InitChannel(queueueName string) error {
	var err error
	b.Channel, err = b.Conn.Channel()
	if err != nil {
		return err
	}

	q, err := b.Channel.QueueDeclare(
		queueueName, // Nombre de la cola
		true,        // Duradera
		false,       // Autoeliminar si no hay consumidores
		false,       // Exclusiva
		false,       // No espera confirmación del servidor
		nil,         // Argumentos adicionales
	)
	if err != nil {
		return err
	}

	b.QueueueName = queueueName

	fmt.Println("Cola declarada:", q.Name)

	return nil
}

func (b *RabbitMQBroker) Publish(message string) error {
	err := b.Channel.Publish(
		"",            // Exchange
		b.QueueueName, // Routing key (nombre de la cola)
		false,         // Obligatorio
		false,         // Inmediato
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return err
	}

	fmt.Println("Mensaje enviado")

	return nil
}
