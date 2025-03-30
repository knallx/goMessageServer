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
   
	 // Process and use the data
	 fmt.Printf("Received: %s\n", buffer[:n])
	}
   }

func sendData(conn net.Conn) {
	var USRinput string
	USRinput = "SYN"
	
	// Gather input from terminal
	for USRinput != "" {
		fmt.Printf("[+]:")
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
    // Connect to the server
	ipInput := ""
	fmt.Print("Enter your IP to connect to\n[+]:")
	fmt.Scan(&ipInput)
	ipInput = ipInput + ":"
	portInput := ""
	fmt.Print("Enter port to connect to\n[+]:")
	fmt.Scan(&portInput)
	ipInput = ipInput + portInput

	// Dial up the server
    conn, err := net.Dial("tcp", ipInput)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer conn.Close()

	// Runs these functions in go routines in the background
	go sendData(conn)
	go handleClient(conn)

	// Just keeps this alive until you ctrl+c
	select {}
	
}
    
