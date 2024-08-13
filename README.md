# Learning project - Real-time sensor data streaming over a custom TCP protocol

## The project

This is a learning project to understand how to stream sensor data over a custom TCP protocol. The project is divided into two parts: the **server** and the **client**. The server is responsible for receiving sensor data from multiple sensors and logging it first to the console and then to a file. The client is responsible for sending sensor data to the server. (The client is more like a dummy sensor for now)


## The protocol
The reason why I chose to send the data using a custom protocol is to understand how to create a protocol from scratch. The protocol is a simple one and is as follows:

### Message format
- Sensor Data format
    - Message Type (1 byte) - `0x01` (unused for now)
    - Sensor ID (1 byte) - `0x01` to `0xFF`
    - Timestamp (4 bytes) - Unix timestamp in seconds
    - Data Length (4 bytes) - Length of the data
    - Data (Variable length) - Sensor data

### Message types
- `0x01` - Sensor Data
- More to be added...


## Features
- [x] Server listens for incoming connections
- [x] The sensor client connects to the server
- [x] Server receives sensor data from the client
- [x] Server can handle multiple clients
- [x] Server logs the sensor data to the console
- [x] The client retries to connect to the server if the connection is lost
- [ ] Server logs the sensor data to a file

## How to run?
1. Clone the repository
2. Run the server
    - Run `go run server/server.go`
3. Run the sensor client
    - Run `go run client/sensor.go`