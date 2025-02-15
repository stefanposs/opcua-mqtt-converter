# OPC UA to MQTT Converter - Concept

## 1. Project Goal

The goal of this project is to develop a robust, flexible, and configurable converter that retrieves data from an OPC UA server and forwards it in real-time to an MQTT broker. The converter will be developed in Go (Golang) and offer high performance and scalability. It requires a buffer to store intermediate states in case of connection failures or container crashes and resend them when needed - Retention Policy.

## 2. Use Cases

- **Industry 4.0**: Machines send sensor data (temperature, pressure, humidity) to a central dashboard.
- **Building Automation**: OPC UA-based devices in smart buildings provide data to MQTT-based control systems.
- **IoT Gateways**: Conversion of machine data for cloud-based IoT platforms.

## 3. Architecture

### 3.1 Components

- **OPC UA Client**: Connects to the OPC UA server, subscribes to nodes, or polls data at intervals.
- **MQTT Publisher**: Sends the processed data to an MQTT broker.
- **Configuration Manager**: Loads the `config.json` for flexible adjustments without code changes.
- **Logger & Monitoring**: Logs errors, performance data, and connection status.
- **Storage Manager**: Buffers data during connection issues and ensures retention policy.

### 3.2 Data Flow

1. **Initialization**: Configuration file is loaded.
2. **OPC UA Connection**: Client connects to the OPC UA server.
3. **Data Retrieval**: Data is regularly polled or received via subscriptions.
4. **Data Processing**: Optional mapping, transformation, or filtering.
5. **MQTT Publishing**: Data is sent to the MQTT broker.

## 4. Configuration File (`config.json`)

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

## Technology Stack

- **Programming Language**: Go (Golang)
- **Libraries**:
  Preferably use native Go libraries to avoid dependency issues.
  - `github.com/gopcua/opcua` (for OPC UA communication)
  - `github.com/eclipse/paho.mqtt.golang` (for MQTT)

## Security Aspects

- **OPC UA Security**: Support for encryption (Basic256Sha256) and user authentication.
- **MQTT Security**: TLS/SSL encryption and authentication with username/password.
- **Access Control**: Configurable whitelists for OPC UA nodes and MQTT topics.

## Error Handling and Resilience

- Automatic reconnection attempts on connection loss
- Message buffering during MQTT outages
- Detailed logging for debugging and monitoring

## Deployment & Operation

- **Containerization**: Deployment as a Docker container
- **Cloud-Ready**: Usable in Kubernetes clusters
- **Monitoring**: Integration with Prometheus and Grafana for metrics

## Roadmap

### MVP (Minimal Viable Product)

- OPC UA Client (Polling & Subscription)
- MQTT Publisher
- JSON-based configuration

### Extensions

- TLS support for MQTT and OPC UA
- Advanced data processing (e.g., aggregations)
- Web dashboard for status monitoring

## Roles and Responsibilities

- **Product Owner**: Defines requirements, prioritizes features
- **Lead Developer**: Architecture, code review, technical leadership
- **Developer**: Implementation of features and tests
- **DevOps**: CI/CD pipeline, monitoring, deployment

## Risks & Challenges

- Network latency: Optimization of data transmission
- Security requirements: Ensuring end-to-end encryption
- Scalability: Managing high data volumes in industrial applications

## Success Criteria

- **Stability**: Runtime > 99.9% without crashes
- **Performance**: Latency < 100 ms between OPC UA and MQTT
- **Flexibility**: Configuration without code changes

## Design Sketch

### Component Diagram

```plaintext
+-------------------+       +-------------------+       +-------------------+
|                   |       |                   |       |                   |
|   OPC UA Client   +------->   Data Processor  +------->   MQTT Publisher  |
|                   |       |                   |       |                   |
+-------------------+       +-------------------+       +-------------------+
        ^                           |                           |
        |                           v                           v
+-------------------+       +-------------------+       +-------------------+
|                   |       |                   |       |                   |
| Config Manager    |       | Storage Manager   |       | Logger & Monitor  |
|                   |       |                   |       |                   |
+-------------------+       +-------------------+       +-------------------+
```

### Data Flow Diagram

1. **Initialization**: Config Manager loads the configuration file.
2. **OPC UA Connection**: OPC UA Client connects to the OPC UA server.
3. **Data Retrieval**: Data is regularly polled or received via subscriptions.
4. **Data Processing**: Data Processor performs optional mapping, transformation, or filtering.
5. **MQTT Publishing**: MQTT Publisher sends the processed data to the MQTT broker.
6. **Error Handling**: Storage Manager buffers data during connection issues.
7. **Logging & Monitoring**: Logger & Monitor logs errors, performance data, and connection status.

This detailed description and design should help you develop the code and ensure the system is maintainable, interchangeable, and scalable.
