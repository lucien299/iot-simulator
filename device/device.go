package device

import (
	"iot-simulator/config"
	"iot-simulator/my_mqtt"
)

//公共基础设施

// Device 接口
type Device interface {
	GenerateData() ([]byte, error)
	Run()
}

type BaseDevice struct {
	ID     string
	Topic  string
	Client *my_mqtt.MQTTClient
	Config *config.Config
}

func (b *BaseDevice) Run() {
	// 每个设备的通用逻辑可以放在这里，例如启动数据生成和定期发送
	// 具体的设备类型会实现自己的 GenerateData 方法
}
