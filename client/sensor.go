package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// Constant for the Sensor ID
const SENSOR_ID = 1234

func sendData(conn net.Conn) {
	for {
		// Simulate sensor data
		timestamp := time.Now().Unix()

		// Simulate sensor data, e.g., temperature
		data := []byte(fmt.Sprintf("%d", 20+time.Now().Unix()%10))
		dataLen := len(data)

		// Create message based on custom protocol
		msg := make([]byte, 1+4+8+4+dataLen) // Byte slice with variable length based on actual data

		// Fill in message fields
		msg[0] = 0x01                                               // Message type (unused for now)
		binary.LittleEndian.PutUint32(msg[1:5], SENSOR_ID)          // Sensor ID
		binary.LittleEndian.PutUint64(msg[5:13], uint64(timestamp)) // Timestamp
		binary.LittleEndian.PutUint32(msg[13:17], uint32(dataLen))  // Data length
		copy(msg[17:], data)                                        // Actual data

		fmt.Println("Sending sensor data:", string(data))

		// Send message to server
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("Error sending data:", err)
			fmt.Println("Connection lost. Attempting to reconnect...")
			return // Exit the loop and function if there is an error
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	for {
		// Attempt to connect to the server
		conn, err := net.Dial("tcp", "localhost:9000")
		if err != nil {
			fmt.Println("Error connecting to server:", err)
			fmt.Println("Retrying in 5 seconds...")
			time.Sleep(5 * time.Second) // Wait before retrying
			continue
		}

		// Successfully connected
		fmt.Println("Connected to server:", conn.RemoteAddr())

		// Send sensor data until the connection is lost
		sendData(conn)

		// Connection lost, so close the connection
		conn.Close()

		// Optionally, wait before attempting to reconnect
		time.Sleep(5 * time.Second)
	}
}
