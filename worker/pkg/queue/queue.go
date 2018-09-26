package queue

import (
	"log"

	nats "github.com/nats-io/go-nats"
)

type NatsConnection struct {
	connection *nats.Conn
	queue      string
}

type Job interface {
	Run() error
}

// Connect creates a new connection to Nats.io
func Connect(urls string, queue string) (*NatsConnection, error) {
	nc, err := nats.Connect(urls)
	if err != nil {
		return nil, err
	}
	return &NatsConnection{
		connection: nc,
		queue:      queue,
	}, nil
}

func (c *NatsConnection) GetIncomingMessages(ch chan []byte) {
	c.connection.Subscribe(c.queue, func(m *nats.Msg) {
		log.Printf("Received a message: %s\n", string(m.Data))
		ch <- m.Data
	})
}

func (c *NatsConnection) Close() error {
	c.connection.Close()
	return nil
}
