package main

import (
	"os"
	"strings"
)

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
