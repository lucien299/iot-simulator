package main

import (
	"iot-simulator/config"
	"iot-simulator/mqtt"
	"iot-simulator/simulator"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	client, err := mqtt.NewClient(mqtt.MQTTOptions{
		Broker:   cfg.MQTT.Broker,
		Username: cfg.MQTT.Username,
		Password: cfg.MQTT.Password,
		Topic:    cfg.Topic.TemperatureHumiditySensorTopic,
		QoS:      0,
		ClientID: "",
	})
	if err != nil {
		log.Fatalf("MQTT connection error: %v", err)
	}

	defer client.Disconnect()

	log.Println("Starting simulation...")
	simulator.StartSimulation(cfg, *client)
	log.Println("Simulation completed.")
}
