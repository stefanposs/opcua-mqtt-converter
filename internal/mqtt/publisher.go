package mqtt

import (
	"github.com/eclipse/paho.mqtt.golang"
)

type Publisher struct {
	client mqtt.Client
	topic  string
	qos    byte
	retain bool
}

func NewPublisher(broker, topic string, qos byte, retain bool) *Publisher {
	opts := mqtt.NewClientOptions().AddBroker(broker)
	client := mqtt.NewClient(opts)
	return &Publisher{client: client, topic: topic, qos: qos, retain: retain}
}

func (p *Publisher) Connect() error {
	token := p.client.Connect()
	token.Wait()
	return token.Error()
}

func (p *Publisher) Publish(message string) error {
	token := p.client.Publish(p.topic, p.qos, p.retain, message)
	token.Wait()
	return token.Error()
}

func (p *Publisher) Disconnect() {
	p.client.Disconnect(250)
}
