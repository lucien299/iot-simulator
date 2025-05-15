package main

import (
	"iot-simulator/config"
	"iot-simulator/my_mqtt"
	"iot-simulator/simulator"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	client, err := my_mqtt.NewClient(my_mqtt.MQTTOptions{
		Broker:   cfg.MQTT.Broker,
		Username: cfg.MQTT.Username,
		Password: cfg.MQTT.Password,
		Topic:    cfg.Topic.TemperatureHumiditySensorTopic,
		QoS:      1,
		ClientID: "go_sim",
	})
	if err != nil {
		log.Fatalf("MQTT connection error: %v", err)
	}

	defer client.Disconnect()

	log.Println("Starting simulation...")
	simulator.StartSimulation(cfg, *client)
	log.Println("Simulation completed.")
}
