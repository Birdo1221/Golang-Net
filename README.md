# 🚀 Go Botnet Command and Control (C&C) Server

## ⚠️ Project Status: Deprecated
This project is no longer maintained. For the latest version, please visit the [Enhanced Version](https://github.com/Birdo1221/Better-Go-Cnc/).

## 📋 Overview
A simple Command and Control (C&C) server implementation in Go, created for educational purposes to understand network programming and client-server architecture.

## 🚦 Prerequisites
- Go programming language installed on your system

### Installation Guide

#### Windows
Download and install Go from [official website](https://go.dev/dl/)

#### Linux
Choose your package manager:
```bash
# Debian/Ubuntu
apt install golang

# RHEL/CentOS
yum install golang

# Fedora
dnf install golang
```

## ⚙️ Configuration

### Bot Configuration
Edit these values in `bot.go`:
```go
const (
    cncServerIP   = "127.0.0.1"  // Your CNC server IP
    cncServerPort = ":5000"      // Your CNC server port
)

const (
    botPort = ":5000"  // Bot listening port
)
```

### Server Configuration
Edit these values in `main.go`:
```go
cncServerIP := "127.0.0.1"    // Server binding IP
cncServerPort := "6000"       // User login port
```

## 🔧 Installation

1. Clone the repository
2. Navigate to the project directory
3. Build and run the components:

```bash
# Build components
go build main.go
go build bot.go

# Run components
go run main.go
go run bot.go
```

## 🚀 Usage

### Starting the Server
```bash
./main
```

### User Authentication
Create a `Users.txt` file with credentials in the following format:
```
Admin:Passwd
```

### Available Commands
- `bots`: Display number of connected bots
- `clear`: Clear bot terminal screens
- `logout`: Disconnect from C&C server

### Connecting to C&C
Use MobaXterm or PuTTY to connect:
1. Enter server IP address
2. Use configured bot port
3. Login with credentials from Users.txt

## ⚠️ Security Notice
This project is:
- **Educational Only**: Created for learning network programming
- **Not Production-Ready**: Lacks security features
- **Local Testing Only**: Not suitable for public networks
- **Proof of Concept**: Demonstrates basic C&C architecture

## ⚖️ Legal Disclaimer
This code is for educational purposes only. Usage for malicious purposes is strictly prohibited. Users must:
- Understand local cybersecurity laws
- Obtain proper authorization for testing
- Use only in controlled environments
- Accept full responsibility for usage

## 📝 License
This project is licensed under the MIT License - see the LICENSE file for details.
