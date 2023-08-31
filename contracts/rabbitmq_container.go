package contracts

import (
	"context"
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQContainerOptions struct {
	Host        string
	VirtualHost string
	Ports       []string
	HostPort    int
	HttpPort    int
	UserName    string
	Password    string
	ImageName   string
	Name        string
	Tag         string
}

type RabbitMQContainer interface {
	Start(ctx context.Context,
		t *testing.T) (*amqp.Channel, error)

	Cleanup(ctx context.Context) error
}
