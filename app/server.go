package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
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
                postFile(filePath,body, conn)
            }
		} else {
			_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
			if err != nil {
				fmt.Println("Error writing on the connection: ", err.Error())
			}
		}
	}
}

func extractPath(req string) string {
	var path string
	start := strings.Index(req, " ") + 1
	end := strings.Index(req[start:], " ") + start
	if start > 0 && end > start {
		path = req[start:end]
	}
	return path
}

func extractMethod(req string) string {
	end := strings.Index(req, " ")
	return req[:end]
}

func extractUserAgent(req []byte) string {
	lines := strings.Split(string(req), "\r\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "User-Agent: ") {
			return strings.TrimPrefix(line, "User-Agent: ")
		}
	}
	return ""
}

func extractPostBody(req []byte) string {
    lines := strings.Split(string(req), "\r\n")
    for i := len(lines) - 1; i >= 0; i-- {
        line := strings.TrimRight(lines[i], "\x00")
        line = strings.TrimSpace(line)
        if line != "" {
            return line
        }
    }
    return ""
}


func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
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

func postFile(filePath string,body string, conn net.Conn) {
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

