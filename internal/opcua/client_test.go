package opcua

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient("opc.tcp://localhost:4840")
	assert.NotNil(t, client)
}

func TestClientConnect(t *testing.T) {
	client := NewClient("opc.tcp://localhost:4840")
	ctx := context.Background()
	err := client.Connect(ctx)
	assert.Nil(t, err)
	client.Disconnect()
}
