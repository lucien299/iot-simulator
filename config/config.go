package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type MQTTConfig struct {
	Broker   string `yaml:"broker"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	QoS      int    `yaml:"qos"`
}

type SimulationConfig struct {
	DeviceCount int           `yaml:"device_count"`
	Frequency   time.Duration `yaml:"frequency"`
	Duration    time.Duration `yaml:"duration"`
}

type Topic struct {
	TemperatureHumiditySensorTopic string `yaml:"temperature_humidity_sensor"`
}
type Config struct {
	MQTT       MQTTConfig       `yaml:"mqtt"`
	Simulation SimulationConfig `yaml:"simulation"`
	Topic      Topic            `yaml:"topics"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
