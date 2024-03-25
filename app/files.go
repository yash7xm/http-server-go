package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"path/filepath"
	"strings"
)

func handleFileRequest(conn net.Conn, path string, method string, req []byte) {
	fileName := strings.TrimPrefix(path, "/files/")
	filePath := filepath.Join(directoryPath, fileName)
	if method == "GET" {
		getFile(filePath, conn)
	} else if method == "POST" {
		body := extractPostBody(req)
		postFile(filePath, body, conn)
	}
}

func getFile(filePath string, conn net.Conn) {
	if fileExists(filePath) {
		fileContents, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			_, err := conn.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n"))
			if err != nil {
				fmt.Println("Error writing on the connection: ", err.Error())
			}
			return
		}
		response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n%s", len(fileContents), fileContents)
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Error writing on the connection: ", err.Error())
		}
	} else {
		_, err := conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		if err != nil {
			fmt.Println("Error writing on the connection: ", err.Error())
		}
	}
}

func postFile(filePath string, body string, conn net.Conn) {
	err := ioutil.WriteFile(filePath, []byte(body), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		_, err := conn.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n"))
		if err != nil {
			fmt.Println("Error writing on the connection: ", err.Error())
		}
		return
	}
	response := "HTTP/1.1 201 Created\r\n\r\n"
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing on the connection: ", err.Error())
	}
}
