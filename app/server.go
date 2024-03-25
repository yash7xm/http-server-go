package main

import (
	"fmt"
	"net"
	"os"
)

func main() {	
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	fmt.Println("server started on port 4221")
	
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	req := make([] byte, 1024)
	_, err = conn.Read(req)
	if err != nil {
		fmt.Println("Error reading the connection: ", err.Error())
	}

	_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	if err != nil {
		fmt.Println("Error writing on the connection: ", err.Error())
	}
}
