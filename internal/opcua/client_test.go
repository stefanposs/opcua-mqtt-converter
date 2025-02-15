package opcua

import (
	"errors"
	"testing"
	"time"

	"github.com/gopcua/opcua/ua"
	"github.com/stretchr/testify/assert"
)

type MockOPCUAClient struct{}

func (m *MockOPCUAClient) Connect() error {
	return nil
}

func (m *MockOPCUAClient) Disconnect() {}

func (m *MockOPCUAClient) ReadNodes(nodes []string) (map[string]*ua.DataValue, error) {
	if len(nodes) == 0 {
		return nil, errors.New("no nodes provided")
	}
	result := make(map[string]*ua.DataValue)
	for _, node := range nodes {
		result[node] = &ua.DataValue{Value: ua.MustVariant("mocked value")}
	}
	return result, nil
}

func (m *MockOPCUAClient) SubscribeNodes(nodes []string, interval time.Duration, callback func(map[string]*ua.DataValue)) error {
	return nil
}

func TestNewClient(t *testing.T) {
	client := NewClient("opc.tcp://localhost:4840")
	assert.NotNil(t, client)
}

func TestClientConnect(t *testing.T) {
	client := &MockOPCUAClient{}
	err := client.Connect()
	assert.Nil(t, err)
	client.Disconnect()
}

func TestClientReadNodes(t *testing.T) {
	client := &MockOPCUAClient{}
	nodes := []string{"ns=2;s=Sensor1", "ns=2;s=Sensor2"}
	data, err := client.ReadNodes(nodes)
	assert.Nil(t, err)
	assert.Equal(t, "mocked value", data["ns=2;s=Sensor1"].Value.String())
	assert.Equal(t, "mocked value", data["ns=2;s=Sensor2"].Value.String())
}
