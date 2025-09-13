package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Start listening on 9092")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Connection accepted on 9092")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	received := bytes.Buffer{}
	buff := make([]byte, 1024)

	fmt.Println("Connection read on 9092")

	_, err := conn.Read(buff)
	if err != nil && err == io.EOF {
		fmt.Printf("failed to read from connection: %v", err)
	}

	fmt.Println("Connection read done on 9092")

	received.Write(buff)
	// request := request_decoder.DecodeRequest(received.Bytes())
	// response := handler.HandlerRequest(request)

	conn.Write([]byte{0, 0, 0, 0, 0, 0, 0, 7})
}
