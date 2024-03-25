package main

import (
	"fmt"
	"net"
	"path/filepath"
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
			randomString := strings.TrimPrefix(path, "/echo/")
			response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(randomString), randomString)
			_, err := conn.Write([]byte(response))
			if err != nil {
				fmt.Println("Error writing on the connection: ", err.Error())
			}
		} else if path == "/user-agent" {
			ua := extractUserAgent(req[:n])
			response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(ua), ua)
			_, err := conn.Write([]byte(response))
			if err != nil {
				fmt.Println("Error writing on the connection: ", err.Error())
			}
		} else if strings.HasPrefix(path, "/files/") {
			fileName := strings.TrimPrefix(path, "/files/")
			filePath := filepath.Join(directoryPath, fileName)
			if method == "GET" {
				getFile(filePath, conn)
			} else if method == "POST" {
				fmt.Println("File content is:- ", string(req[n:]))
				body := extractPostBody(req)
				fmt.Println("Body is :- ", body)
				postFile(filePath, body, conn)
			}
		} else {
			_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
			if err != nil {
				fmt.Println("Error writing on the connection: ", err.Error())
			}
		}
	}
}
