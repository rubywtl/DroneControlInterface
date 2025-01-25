# DroneControlInterface 

This project sets up a WebSocket server using the Gin framework in Go and the Gorilla WebSocket package. It can be used to control a drone in real-time via WebSocket communication between the client and server.

## Prerequisites

Before running the server, make sure you have the following installed on your machine:

### 1. **Go (Golang)**

You need to have Go installed to run the server. You can download it from the official Go website:

- [Download Go](https://go.dev/dl/)

After installation, verify that Go is installed by running:

```bash
go version
```

### 2. Dependencies
The project uses the following Go dependencies:
- Gin for the web framework.
- Gorilla WebSocket for WebSocket handling.
```bash
go get github.com/gin-gonic/gin
go get github.com/gorilla/websocket
```

## Running The Server
### 1. Clone Repo
```bash
git clone https://github.com/your-username/DroneControlInterface.git
cd DroneControlInterface
```
### 2. Run Server
```bash
go run main.go
```
### 3. Testing WebSocket Client
You can test the server using the provided index.html file. Open it in a web browser to test the WebSocket connection.