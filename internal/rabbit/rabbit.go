package rabbit

import (
	"fmt"

	"github.com/labstack/gommon/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"uala.com/core-service/config"
)

type Rabbit struct {
	con *amqp.Connection
}

func NewRabbit(conf *config.Config) *Rabbit {
	fmt.Println("Connecting to RabbitMQ")
	conex := fmt.Sprintf("amqp://%s:%s@%s:%d/", conf.Rabbit.User, conf.Rabbit.Password, conf.Rabbit.Host, conf.Rabbit.Port)
	fmt.Println(conex)
	conn, err := amqp.Dial(conex)

	if err != nil {
		log.Error("Failed to connect to RabbitMQ", err)
	}
	return &Rabbit{con: conn}

}

func (r *Rabbit) GetConnection() *amqp.Connection {
	return r.con
}

func (r *Rabbit) Close() {
	r.con.Close()
}

func (r *Rabbit) GetChannel() (*amqp.Channel, error) {
	return r.con.Channel()
}

func (r *Rabbit) CreateQueue(chanel *amqp.Channel) (amqp.Queue, error) {
	return chanel.QueueDeclare("main", true, false, false, false, nil)
}

func (r *Rabbit) Publish(chanel *amqp.Channel, queue amqp.Queue, body []byte) error {
	return chanel.Publish("", queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
}
