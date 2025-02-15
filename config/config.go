package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	OPCUA struct {
		Endpoint        string   `json:"endpoint"`
		Nodes           []string `json:"nodes"`
		PollingInterval int      `json:"polling_interval"`
	} `json:"opcua"`
	MQTT struct {
		Broker string `json:"broker"`
		Topic  string `json:"topic"`
		QoS    byte   `json:"qos"`
		Retain bool   `json:"retain"`
	} `json:"mqtt"`
	Storage struct {
		Path string `json:"path"`
	} `json:"storage"`
	RetentionPolicy struct {
		Enabled bool `json:"enabled"`
		MaxSize int  `json:"max_size"`
		MaxAge  int  `json:"max_age"`
	} `json:"retention_policy"`
	Logging struct {
		Level string `json:"level"`
		File  string `json:"file"`
	} `json:"logging"`
}

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
