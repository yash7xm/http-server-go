# HTTP Server

## Overview

This is a simple HTTP server written in Go that provides various functionalities such as echoing a string, retrieving user agent information, and handling file operations like fetching and saving files. It supports both GET and POST methods.

## Features

- Echo functionality: Echoes back any string provided in the URL.
- Header: Retrieves and returns the header information of the client.
- File operations:
  - GET: Fetches the contents of a file from the server directory.
  - POST: Saves the contents of a file sent in the request body to the server directory.

## Setup

1. Clone the repository: `git clone https://github.com/yash7xm/http-server-go`
2. Navigate to the project directory: `cd http-server-go`
3. Run by using: `./my_server.sh`

## Usage

### Echo

To echo a string, append it to the URL after `/echo/`: GET /echo/<your_string>

Example: 
`GET /echo/hello_world`
Response: 
```
HTTP/1.1 200 OK
Content-Type: text/plain
Content-Length: 11

hello_world
```

### Reading Header

To retrieve the header information: `GET /user-agent`
Response: 
```
HTTP/1.1 200 OK
Content-Type: text/plain
Content-Length: 11

<header_info>
```

### File Operations

#### GET

To fetch the contents of a file: `GET /files/<filename>`

Response: 
```
HTTP/1.1 200 OK
Content-Type: application/octet-stream
Content-Length: <file_size>

<file_contents>
```

#### POST

To save contents to a file: `POST /files/<filename>`

Request Body: 
```
Contents of the file
Response: HTTP/1.1 201 Created
```

## Dependencies

- Go 1.13 or higher

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.
