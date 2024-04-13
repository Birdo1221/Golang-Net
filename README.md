:rocket: Simple Botnet Command and Control (C&C) Server in Go.

Ive Updated This Project For Both the Server And Client 
On My New Repo : 
```[ Click For Better Varient ](https://github.com/Birdo1221/Better-Go-Cnc/edit/main/README.md)```

:white_check_mark: Getting Started


Make sure you have Go installed on your system.
You can download and install it from the official site

For Windows
```
https://go.dev/dl/
```
Or for Linux Based OS'S
```
apt install golang
yum install golang
dnf install golang
```

:wrench:   You may need to configure the files for there IP's or Port Values 

Edit Line 10-11 to change the IP to connect to and the Port for the Bot connection 

```
const (
	cncServerIP   = "127.0.0.1" // Replace with your CNC server IP
	cncServerPort = ":5000"     // Replace with your CNC server port
)
```

Edit Line 14 to change the Bot Port for the Port the Bot / Client Connects at.
```
const (
	botPort         = ":5000"   // Bot will listen on port 5000
)
```

Edit Line 218-219 for the server IP to bind to and Port for the User Login on the server
```
   cncServerIP := "127.0.0.1"
   cncServerPort := "6000"
```


To compile the C&C server and the bot server, follow these steps:

Clone this repository, Change folder dir to the main program file and using the Go compiler
```
go build main.go

go build bot.go

go run main.go
go run bot.go    
```

:gear:  Running the C&C Server

Start the C&C server on port 1338 [default port].
```
./main
```

After Start the Bot / Client  on a virtual machine or on your local machine to test it out, monitor you network to see the packet to the [ Client <--> Bot ] or [ Bot <--> Client ] , i wouldnt recommend using on public foward facing networks as it has no security implementations or anything to save you.
```
./main
```

:speech_balloon:   The C&C server will start listening for incoming connections from Clients on port 1338.
You can change via the SRC.
Login info is on a text file in this format. [Users.txt]
```
Admin:Passwd
```
Once connected to the C&C server, the administrator can send commands.

 
:bulb:   Example Commands
```
bots: Get the number of connected bots.

clear: Clear the screen of the connected bot terminals.

logout: Log out from the C&C server.
```

To connect To the cnc I would recommend using mobaxterm or putty use the ip the server is running on and the botport and then you will be a login menu for the cnc

❗ Remember that this is a POC / concept

:lock:   Security Considerations

This project is intended for educational and learning purposes only. The code provided here is a very simplified example of a botnet architecture. Remember that using a botnet for any malicious purposes is strictly prohibited and against ethical guidelines. This project is intended for educational and learning purposes only.

The code does not include advanced security measures, and it is not suitable for Public use. If you plan to build a real botnet or similar systems, ensure that you understand the legal implications of you actions and the importance of securing  system's to prevent unauthorized access.

:handshake:    Contributions are welcome! If you find any issues or have improvements to suggest, feel free to create a pull request or just reupload, no need to reference me for the SRC.

:page_with_curl:    This project is licensed under the MIT License - see the LICENSE file for details.

