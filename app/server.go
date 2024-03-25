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

	req := make([]byte, 1024)
	n, err := conn.Read(req)
	if err != nil {
		fmt.Println("Error reading the connection: ", err.Error())
	}

	fmt.Println("req read from conn:- ", string(req[:n]))

	path := extractPath(string(req[:n]))

	fmt.Println("Path from the req is:- ", path)

	if path == "/" {
		_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		if err != nil {
			fmt.Println("Error writing on the connection: ", err.Error())
		}
	} else {
		_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		if err != nil {
			fmt.Println("Error writing on the connection: ", err.Error())
		}
	}

}

func extractPath(req string) string {
	var path string
	if req[4] == '/' {
		for i := 4; i < len(req); i++ {
			if string(req[i]) != " " {
				path += string(req[i])
			} else {
				break
			}
		}
	}

	return path
}
