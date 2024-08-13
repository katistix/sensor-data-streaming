package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connection established:", conn.RemoteAddr())

	for {
		msg := make([]byte, 1024) // Adjust buffer size as needed
		n, err := conn.Read(msg)
		if err != nil {
			fmt.Println("Error reading data:", err)
			return
		}

		if n > 0 {
			// Process data based on custom protocol
			messageType := msg[0]
			sensorID := binary.LittleEndian.Uint32(msg[1:5])
			timestamp := binary.LittleEndian.Uint64(msg[5:13])
			dataLen := binary.LittleEndian.Uint32(msg[13:17])
			data := msg[17 : 17+dataLen]

			// Print the data
			fmt.Printf("Message Type: %d\n", messageType)
			fmt.Printf("Sensor ID: %d, Timestamp: %d, Data: %s\n", sensorID, timestamp, string(data))
			fmt.Printf("----------------------------------------\n")
		}
	}
}
func main() {
	// Start server
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on port 9000")

	// Accept incoming connections
	// And handle them in a separate goroutine
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
