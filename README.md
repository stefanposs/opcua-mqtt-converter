# OPC UA to MQTT Converter
[![Build Status](https://img.shields.io/github/actions/workflow/status/stefanposs/opcua-mqtt-converter/coveralls.yml?branch=main)](https://github.com/stefanposs/opcua-mqtt-converter/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/stefanposs/opcua-mqtt-converter)](https://goreportcard.com/report/github.com/stefanposs/opcua-mqtt-converter)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Website](https://img.shields.io/badge/Website-Visit-blue)](https://stefanposs.de)
[![Coverage Status](https://coveralls.io/repos/github/stefanposs/opcua-mqtt-converter/badge.svg?branch=main)](https://coveralls.io/github/stefanposs/opcua-mqtt-converter?branch=main)
[![GoDoc](https://godoc.org/github.com/stefanposs/opcua-mqtt-converter?status.svg)](https://godoc.org/github.com/stefanposs/opcua-mqtt-converter)
[![Project Status](https://img.shields.io/badge/status-draft-orange)](https://github.com/stefanposs/opcua-mqtt-converter)


## Project Overview

This project is an OPC UA to MQTT converter that retrieves data from an OPC UA server and forwards it in real-time to an MQTT broker. The converter is written in Go (Golang) and offers high performance and scalability. It includes a buffer to temporarily store data in case of connection issues and send it later.

**<span style="color:red">Note:</span>** <span style="color:red">This project is currently in draft status and is not yet stable.</span>


## Use Cases

- **Industry 4.0**: Machines send sensor data (e.g., temperature, pressure) to a central dashboard.
- **Building Automation**: OPC UA-based devices in smart buildings provide data to MQTT-based control systems.
- **IoT Gateways**: Conversion of machine data for cloud-based IoT platforms.

## Architecture

### Components

- **OPC UA Client**: Connects to the OPC UA server and subscribes to nodes or polls data at intervals.
- **MQTT Publisher**: Sends the processed data to an MQTT broker.
- **Configuration Manager**: Loads the `config.json` for flexible adjustments without code changes.
- **Logger & Monitoring**: Logs errors, performance data, and connection status.
- **Storage Manager**: Buffers data during connection issues and ensures retention policy.

### Data Flow

1. **Initialization**: Configuration file is loaded.
2. **OPC UA Connection**: Client connects to the OPC UA server.
3. **Data Retrieval**: Data is regularly polled or obtained via subscriptions.
4. **Data Processing**: Optional mapping, transformation, or filtering.
5. **MQTT Publishing**: Data is sent to the MQTT broker.

## Configuration File (`config.json`)

```json
{
    "opcua": {
        "endpoint": "opc.tcp://localhost:4840",
        "nodes": ["ns=2;s=Sensor1", "ns=2;s=Sensor2"],
        "polling_interval": 5000
    },
    "mqtt": {
        "broker": "tcp://localhost:1883",
        "topic": "sensors/data",
        "qos": 1,
        "retain": false
    },
    "storage": {
        "path": "data"
    },
    "retention_policy": {
        "enabled": true,
        "max_size": 1000,
        "max_age": 3600
    },
    "logging": {
        "level": "info",
        "file": "converter.log"
    }
}
```

### Parameter Explanation:

- `opcua.endpoint`: OPC UA server address
- `nodes`: List of nodes to monitor
- `polling_interval`: Polling interval in milliseconds
- `mqtt.broker`: MQTT broker address
- `topic`: MQTT topic for publishing
- `qos`: Quality of Service level
- `logging`: Logging settings

## Installation

1. **Clone the repository:**
     ```sh
     git clone https://github.com/stefanposs/opcua-mqtt-converter.git
     cd opcua-mqtt-converter
     ```

2. **Install dependencies:**
     ```sh
     go mod tidy
     ```

3. **Adjust configuration file:**
     Edit the `config.json` file according to your environment.

## Usage

1. **Start the converter:**
     ```sh
     go run cmd/main.go
     ```

2. **Logs and data:**
     Check the logs and stored data in the specified directory.

## Tests

Run the tests with the following command:
```sh
go test ./...
```

License
This project is licensed under the MIT License. For more information, see the [License](LICENSE.md) file.

Contribution
Contributions are welcome! Please open an issue or a pull request for suggestions and improvements.