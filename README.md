Simple Botnet Command and Control (C&C) Server

This project demonstrates a simple Botnet Command and Control (C&C) server and bot server written in Go.

:rocket: Getting Started
:white_check_mark: Prerequisites / Requirements 
Go (Golang) - Make sure you have Go installed on your system. You can download and install it from the official 

Go website: https://golang.org/

:wrench: Compiling
To compile the C&C server and the bot server, follow these steps:

Clone this repository to your local machine.

Change into the project directory.
Compile the main program files using the Go compiler:

$ go build main.go

:computer: Usage >

:gear: Running the C&C Server

Before running the C&C server, ensure that you have the necessary user credentials in the users.txt file. Each line should contain a username and password in the format: username:password.

Start the C&C server on port 1338:

$ ./main

The C&C server will start listening for incoming connections from administrators (clients).
:robot: Connecting Bots
On the bots' machines, compile the bot.go file:

$ go build bot.go
Run the bot on each bot machine, specifying the CNC server IP and port:

Edit line 10 - 11 on bot.go to change the ip / port it directs / connects to.
The bot will attempt to connect to the CNC server at the specified IP and port.

:speech_balloon: Sending Commands
Once the bot connects to the C&C server, the administrator can send commands to the connected bots.

:bulb: Example Commands

bots: Get the number of connected bots.

clear: Clear the screen of the connected bot terminals.

logout: Log out from the C&C server.

Any other command will be broadcasted to all connected bots.

To connect To the cnc I would recommend using mobaxterm or putty use the ip the server is running on and the botport and then you will be a login menu for the cnc

❗ Remember that this is a POC / concept

:lock: Security Considerations
This project is intended for educational and learning purposes only. The code provided here is a very simplified example of a botnet architecture. In real-world scenarios, using such code for malicious activities is illegal and unethical.

The code does not include advanced security measures, and it is not suitable for production use. If you plan to build a real botnet or similar systems, ensure that you understand the legal implications and the importance of securing the system to prevent unauthorized access.

:handshake: Contributing
Contributions are welcome! If you find any issues or have improvements to suggest, feel free to create a pull request.

:page_with_curl: License
This project is licensed under the MIT License - see the LICENSE file for details.

Please make sure to include any relevant images (botnet.jpg and botnet_architecture.png) and the actual users.txt file (with the appropriate user credentials) in your repository for users to test the project.

Remember that using a botnet for any malicious purposes is strictly prohibited and against ethical guidelines. This project is intended for educational and learning purposes only.
