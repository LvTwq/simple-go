package mqtt_client

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
	"simple-go/configs"
	"sync"
	"time"
)

var Client *MqttClient

type MqttClient struct {
	client      MQTT.Client
	opts        *MQTT.ClientOptions
	subs        map[string]byte   // 订阅记录：主题 -> QoS
	subMutex    sync.Mutex        // 并发安全锁
	messageChan chan MQTT.Message // 消息通道（可选）
}

func InitMqtt() {
	cfg := configs.AppConfig.Mqtt

	opts := MQTT.NewClientOptions()
	opts.AddBroker(cfg.Host)
	opts.SetClientID(cfg.ClientId)
	opts.SetUsername(cfg.UserName)
	opts.SetPassword(cfg.Password)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(30 * time.Second)

	// 此时还未连接
	Client = &MqttClient{
		opts:        opts,
		subs:        make(map[string]byte),
		messageChan: make(chan MQTT.Message, 100),
	}
}

// Connect 连接并订阅主题
func (m *MqttClient) Connect() error {
	m.opts.OnConnect = func(c MQTT.Client) {
		log.Print("[MQTT] Connected!")
		m.subMutex.Lock()
		defer m.subMutex.Unlock()

		/*	todo	for topic, qos := range m.subs {
			token := c.Subscribe(topic, qos, m.messageChan)
			if token.Wait() && token.Error() != nil {
				log.Printf("[MQTT] 重订阅失败 topic=%s: %v", topic, token.Error())
			}
		}*/
	}

	m.opts.OnConnectionLost = func(c MQTT.Client, err error) {
		log.Printf("[MQTT] Connection lost: %v", err)
	}
	m.client = MQTT.NewClient(m.opts)
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("连接失败: %v", token.Error())
	}
	return nil
}
