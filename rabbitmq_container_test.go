package rabbitmq

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	t.Log("running test 1")
	ctx := context.Background()

	rabbitmq, err := NewRabbitMQTestContainers().Start(ctx, t)

	assert.NotNil(t, rabbitmq)
	assert.NoError(t, err)
}

func Test2(t *testing.T) {
	t.Log("running test 2")
	ctx := context.Background()

	rabbitmq, err := NewRabbitMQTestContainers().Start(ctx, t)

	assert.NotNil(t, rabbitmq)
	assert.NoError(t, err)
}

func Test3(t *testing.T) {
	t.Log("running test 3")
	ctx := context.Background()

	rabbitmq, err := NewRabbitMQTestContainers().Start(ctx, t)

	assert.NotNil(t, rabbitmq)
	assert.NoError(t, err)
}

func Test4(t *testing.T) {
	t.Log("running test 4")
	ctx := context.Background()

	rabbitmq, err := NewRabbitMQTestContainers().Start(ctx, t)

	assert.NotNil(t, rabbitmq)
	assert.NoError(t, err)
}
