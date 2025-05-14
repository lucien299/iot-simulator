package device

import (
	"context"
	"encoding/json"
	"fmt"
	"iot-simulator/config"
	"iot-simulator/mqtt"
	"log"
	"math/rand"
	"time"
)

type TemperatureHumidityData struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Timestamp   int64   `json:"timestamp"`
}

type TemperatureHumidityDevice struct {
	BaseDevice
}

func NewTemperatureHumidityDevice(id string, cfg *config.Config, client *mqtt.MQTTClient) *TemperatureHumidityDevice {
	return &TemperatureHumidityDevice{
		BaseDevice: BaseDevice{
			ID:     id,
			Topic:  fmt.Sprintf(cfg.Topic.TemperatureHumiditySensorTopic),
			Config: cfg,
			Client: client,
		},
	}
}

func (t TemperatureHumidityDevice) GenerateData() ([]byte, error) {
	data := TemperatureHumidityData{
		Temperature: 20 + rand.Float64()*10,
		Humidity:    40 + rand.Float64()*20,
		Timestamp:   time.Now().Unix(),
	}
	return json.Marshal(data)
}

func (t TemperatureHumidityDevice) Run() {
	ticker := time.NewTicker(t.Config.Simulation.Frequency)
	defer ticker.Stop()

	ctx, cancelFunc := context.WithTimeout(context.Background(), t.Config.Simulation.Duration)
	defer cancelFunc()

	for {
		select {
		case <-ticker.C:
			data, err := t.GenerateData()
			if err != nil {
				log.Printf("Failed to generate data: %v", err)
				continue
			}
			err = t.Client.Publish(t.Config.Topic.TemperatureHumiditySensorTopic, data)
			if err != nil {
				log.Printf("Device %d Failed to publish data: %v\n", t.ID, err)
			}
			log.Printf("Device %s Published data to topic %s: %s\n", t.ID, t.Topic, string(data))

		case <-ctx.Done():
			log.Printf("Device %s Stopped", t.ID)
			return
		}
	}
}
