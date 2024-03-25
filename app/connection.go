package main

import (
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	req := make([]byte, 1024)
	n, err := conn.Read(req)
	if err != nil {
		fmt.Println("Error reading the connection: ", err.Error())
	}

	fmt.Println("req read from conn:- ", string(req[:n]))

	path := extractPath(string(req[:n]))
	method := extractMethod(string(req[:n]))

	fmt.Println("Method is:- ", method)
	fmt.Println("Path from the req is:- ", path)

	if path == "/" {
		_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		if err != nil {
			fmt.Println("Error writing on the connection: ", err.Error())
		}
	} else {
		if strings.HasPrefix(path, "/echo/") {
			handleEchoRequest(conn, path)
		} else if path == "/user-agent" {
			handleUserAgentRequest(conn, req[:n])
		} else if strings.HasPrefix(path, "/files/") {
			handleFileRequest(conn, path, method, req[:n])
		} else {
			_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
			if err != nil {
				fmt.Println("Error writing on the connection: ", err.Error())
			}
		}
	}
}
