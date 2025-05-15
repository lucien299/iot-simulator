package simulator

import (
	"iot-simulator/config"
	"iot-simulator/device"
	"iot-simulator/my_mqtt"
	"log"
	"strconv"
	"sync"
)

func StartSimulation(cfg *config.Config, client my_mqtt.MQTTClient) {
	log.Printf("Starting simulation with %d devices", cfg.Simulation.DeviceCount)
	var wg sync.WaitGroup

	for i := 0; i < cfg.Simulation.DeviceCount; i++ {
		wg.Add(1)
		go func(id string) {
			log.Printf("Starting device %s", id)
			defer wg.Done()
			simDevice := device.NewTemperatureHumidityDevice(id, cfg, &client)
			simDevice.Run()
		}("sim_temp_" + strconv.Itoa(i))
	}
	wg.Wait()
}
