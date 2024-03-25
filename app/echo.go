// echo.go
package main

import (
	"fmt"
	"net"
	"strings"
)

func handleEchoRequest(conn net.Conn, path string) {
	randomString := strings.TrimPrefix(path, "/echo/")
	response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(randomString), randomString)
	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing on the connection: ", err.Error())
	}
}
