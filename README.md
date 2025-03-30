# goMessageServer
Super basic golang message server

This is a super crappy golang message server. It uses a server and client to communicate messages over the network. 

To build use the following commands: (Within the directory you download the files)
GOOS=windows GOARCH=amd64 go build -o client.exe client.go
GOOS=windows GOARCH=amd64 go build -o server.exe server.go
You can replace GOOS=<os>, to whatever os you are running. Same with architecure (GOARCH). 

To run the program you will need to run the server by either double clicking the file after build. Or running in the command line. For example in linux you would navigate to the same directory as the files (Once you build them). Then use ./server


Again this is a crappy message golang server... I built this to understand how to send data over the network with golang. Plus how to use go routines. Also this code was frankensteined from mutiple sources, which are:
https://pkg.go.dev/net#pkg-overview
https://okanexe.medium.com/the-complete-guide-to-tcp-ip-connections-in-golang-1216dae27b5a
Also A.I to understand some concepts in the code. 

Also if you are a big dummy like me and find it kinda hard to understand this networking process. This may help? Hence may...
1. Server Starts and is listening for clients
2. Client reach outs and dials server for connection
3. They establish a connection between eachother. The good ole (SYN, SYN/ACK, ACK), ta da connection.
4. Now what? Data Sendith... but wait, we need a buffer to store data for the data.
     The client sends data in chunks, which get stored in the buffer. Then gets read by the receiving device.
5. Message sent and read, boom computer magic.

Final notes, I may or may not add to this program later. I would like to see encryption between the server and client. Besides the point, we will see. 
