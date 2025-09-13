package main

import (
	"fmt"
	"net"
)

func main() {
	// Establish a connection (e.g., to a local server)
	conn, err := net.Dial("tcp", "localhost:9092")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to localhost:8080")
}
