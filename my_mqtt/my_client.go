package my_mqtt

import (
	"errors"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

type MQTTClient struct {
	client mqtt.Client
	topic  string
	qos    int
}

type MQTTOptions struct {
	Broker   string
	Username string
	Password string
	Topic    string
	QoS      int
	ClientID string
}

func NewClient(opts MQTTOptions) (*MQTTClient, error) {
	//设置
	clientOpts := mqtt.NewClientOptions().
		AddBroker(opts.Broker).
		SetClientID(opts.ClientID).
		SetUsername(opts.Username).
		SetPassword(opts.Password).
		SetConnectRetry(true).
		SetAutoReconnect(true).SetConnectionLostHandler(func(client mqtt.Client, err error) {
		log.Println("my_mqtt connection lost:", err)
	}).SetOnConnectHandler(func(client mqtt.Client) {
		log.Println("my_mqtt connection success")
	})

	client := mqtt.NewClient(clientOpts)
	//token代表的操作状态
	token := client.Connect()
	if !token.WaitTimeout(5*time.Second) || token.Error() != nil {
		log.Println("my_mqtt connection error: %v", token.Error())
		return nil, errors.New("my_mqtt connection error")
	}

	return &MQTTClient{
		client: client,
		topic:  opts.Topic,
		qos:    opts.QoS,
	}, nil
}

func (m *MQTTClient) Publish(topic string, payload []byte) error {
	token := m.client.Publish(topic, byte(m.qos), false, payload)
	token.Wait()
	return token.Error()
}

func (m *MQTTClient) Disconnect() {
	m.client.Disconnect(250)
}
