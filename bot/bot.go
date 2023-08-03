package main

import (
	"fmt"
	"net"
	"time"
)

const (
	cncServerIP   = "127.0.0.1" // Replace with your CNC server IP
	cncServerPort = ":5000"    // Replace with your CNC server port
)

func main() {
	// Main loop to continuously attempt to connect to the CNC server
	for {
		conn, err := net.Dial("tcp", cncServerIP+cncServerPort)
		if err != nil {
			fmt.Println("Error connecting to CNC server:", err.Error())
			time.Sleep(10 * time.Second) // Wait 10 seconds before attempting to reconnect
			continue
		}

		fmt.Println("Bot connected to the CNC server.")

		// Loop to handle communication with the CNC server
		for {
			buffer := make([]byte, 1024)
			bytesReceived, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Connection closed by the CNC server. Reconnecting...")
				break
			}

			// Process the received data here
			// Example: printing received data
			fmt.Printf("Received: %s", buffer[:bytesReceived])

			// Example: sending a response
			response := "PONG\n"
			conn.Write([]byte(response))
		}

		conn.Close()
	}
}
