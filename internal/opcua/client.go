package opcua

import (
	"context"
	"time"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
)

type Client struct {
	client *opcua.Client
}

func NewClient(endpoint string) *Client {
	opts := []opcua.Option{
		opcua.SecurityMode(ua.MessageSecurityModeNone),
		opcua.SecurityPolicy(ua.SecurityPolicyURINone),
	}
	c := opcua.NewClient(endpoint, opts...)
	return &Client{client: c}
}
func (c *Client) Connect(ctx context.Context) error {
	return c.client.Connect(ctx)
}

func (c *Client) Disconnect() {
	c.client.Close()
}

func (c *Client) ReadNodes(nodes []string) (map[string]*ua.DataValue, error) {
	ids := make([]*ua.NodeID, len(nodes))
	for i, node := range nodes {
		ids[i] = ua.NewStringNodeID(2, node)
	}
	req := &ua.ReadRequest{
		NodesToRead: make([]*ua.ReadValueID, len(ids)),
	}
	for i, id := range ids {
		req.NodesToRead[i] = &ua.ReadValueID{NodeID: id}
	}
	resp, err := c.client.Read(req)
	if err != nil {
		return nil, err
	}
	result := make(map[string]*ua.DataValue)
	for i, res := range resp.Results {
		result[nodes[i]] = res
	}
	return result, nil
}

func (c *Client) SubscribeNodes(nodes []string, interval time.Duration, callback func(map[string]*ua.DataValue)) error {
	// Implement subscription logic
	return nil
}
