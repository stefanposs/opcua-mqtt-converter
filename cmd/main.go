package main

import (
	"context"
	"log"
	"time"

	"github.com/stefanposs/opcua-mqtt-converter/config"
	"github.com/stefanposs/opcua-mqtt-converter/internal/logger"
	"github.com/stefanposs/opcua-mqtt-converter/internal/mqtt"
	"github.com/stefanposs/opcua-mqtt-converter/internal/opcua"
	"github.com/stefanposs/opcua-mqtt-converter/internal/processor"
	"github.com/stefanposs/opcua-mqtt-converter/internal/storage"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	log, err := logger.NewLogger(cfg.Logging.Level, cfg.Logging.File)
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}

	opcuaClient := opcua.NewClient(cfg.OPCUA.Endpoint)
	ctx := context.Background()
	if err := opcuaClient.Connect(ctx); err != nil {
		log.Fatalf("Error connecting to OPC UA server: %v", err)
	}
	defer opcuaClient.Disconnect()

	mqttPublisher := mqtt.NewPublisher(cfg.MQTT.Broker, cfg.MQTT.Topic, cfg.MQTT.QoS, cfg.MQTT.Retain)
	if err := mqttPublisher.Connect(); err != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", err)
	}
	defer mqttPublisher.Disconnect()

	storageManager := storage.NewManager(cfg.Storage.Path)
	processor := processor.NewProcessor()

	ticker := time.NewTicker(time.Duration(cfg.OPCUA.PollingInterval) * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			data, err := opcuaClient.ReadNodes(cfg.OPCUA.Nodes)
			if err != nil {
				log.Errorf("Error reading nodes: %v", err)
				continue
			}

			convertedData := make(map[string]interface{})
			for k, v := range data {
				convertedData[k] = v
			}
			processedData, err := processor.Process(convertedData)
			if err != nil {
				log.Errorf("Error processing data: %v", err)
				continue
			}

			if err := mqttPublisher.Publish(processedData); err != nil {
				log.Errorf("Error publishing data: %v", err)
				continue
			}

			convertedData = make(map[string]interface{})
			for k, v := range data {
				convertedData[k] = v
			}
			if err := storageManager.Save(convertedData); err != nil {
				log.Errorf("Error saving data: %v", err)
			}
		}
	}
}
