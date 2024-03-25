package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

var directoryPath string

func main() {

	flag.StringVar(&directoryPath, "directory", "", "Directory path")
	flag.Parse()

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	fmt.Println("server started on port 4221")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
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
