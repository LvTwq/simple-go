package configs

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	AppName  string         `yaml:"app_name"`
	Port     string         `yaml:"port"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	Mqtt     MqttConfig     `yaml:"mqtt_client"`
}

type DatabaseConfig struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Charset  string `yaml:"charset"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type MqttConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	UserName string `yaml:"username"`
	ClientId string `yaml:"client_id"`
}

var AppConfig Config

func InitConfig() {
	file, err := os.Open("configs/config.yaml")
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Error closing config file: %v", err)
		}
	}(file)

	err1 := yaml.NewDecoder(file).Decode(&AppConfig)
	if err1 != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}
}
