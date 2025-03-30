package main

import (
    "fmt"
    "net"
)

func handleClient(conn net.Conn) {
	defer conn.Close()
   
	// Create a buffer to read data into
	buffer := make([]byte, 1024)
   
	for {
	 // Read data from the client
	 n, err := conn.Read(buffer)
	 if err != nil {
	  fmt.Println("Error:", err)
	  return
	 }
   
	 // Process and use the data (here, we'll just print it)
	 fmt.Printf("Client Sent: %s\n", buffer[:n])
	}
   }

func sendData (conn net.Conn) {
	var USRinput string
	USRinput = "SYN"
	
	// Gather input from terminal
	for USRinput != "" {
		fmt.Print("[+]:")
		fmt.Scan(&USRinput)

		// Send data to the server
		data := []byte(USRinput)
		_, err := conn.Write(data)
		if err != nil {
		fmt.Println("Error:", err)
		return
		}
	}
}

func main() {
    // Listen for incoming connections
	ipInput := ""
	fmt.Print("Enter your IP to host the server\n[+]:")
	fmt.Scan(&ipInput)
	ipInput = ipInput + ":"
	portInput := ""
	fmt.Print("Enter your port to serve on\n[+]:")
	fmt.Scan(&portInput)
	ipInput = ipInput + portInput

	// Start up the listener
    listener, err := net.Listen("tcp", ipInput)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Server is listening on port 8080")

    for {
        // Accept incoming connections
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }

        // Runs these functions in go routines in the background
        go handleClient(conn)
		go sendData(conn)
    }
}

