Simple Botnet Command and Control (C&C) Server in Go.

:rocket: Getting Started

:white_check_mark: 
Make sure you have Go installed on your system.
You can download and install it from the official site
```
https://go.dev/
```
Or for Linux Based OS'S
```
apt install golang
yum install golang
dnf install golang
```

:wrench: 
To compile the C&C server and the bot server, follow these steps:

Clone this repository, Change folder dir to the main program file and using the Go compiler

```
go build main.go
go run main.go   
go build bot.go  
go run bot.go    
```

:gear: Running the C&C Server
StartS the C&C server on port 1338 [default]:
```
./main
```

:speech_balloon: 
The C&C server will start listening for incoming connections from Clients on port 1338.
You can change via the SRC.
Login info is on a text file in this format. [Users.txt]

Username:Passwd

Edit line 10 - 11 on bot.go to change the ip / port it listens / connects to.
The bot will attempt to connect to the server at the specified IP and port.
 
Once connected to the C&C server, the administrator can send commands.

:bulb: Example Commands
```
bots: Get the number of connected bots.

clear: Clear the screen of the connected bot terminals.

logout: Log out from the C&C server.
```

To connect To the cnc I would recommend using mobaxterm or putty use the ip the server is running on and the botport and then you will be a login menu for the cnc

❗ Remember that this is a POC / concept

:lock: Security Considerations

This project is intended for educational and learning purposes only. The code provided here is a very simplified example of a botnet architecture. Remember that using a botnet for any malicious purposes is strictly prohibited and against ethical guidelines. This project is intended for educational and learning purposes only.

The code does not include advanced security measures, and it is not suitable for Public use. If you plan to build a real botnet or similar systems, ensure that you understand the legal implications of you actions and the importance of securing  system's to prevent unauthorized access.

:handshake: 
Contributions are welcome! If you find any issues or have improvements to suggest, feel free to create a pull request or just reupload, no need to reference me for the SRC.

:page_with_curl: 
This project is licensed under the MIT License - see the LICENSE file for details.

