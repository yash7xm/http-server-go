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
