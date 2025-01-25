package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "net/http"
)

// Upgrader to upgrade HTTP connection to WebSocket
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true 
    },
}


func wsHandler(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println("Upgrade failed:", err)
        return
    }
    defer conn.Close()

    log.Println("New WebSocket connection established!")

    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println("Read message failed:", err)
            break
        }

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println("Write message failed:", err)
            break
        }
    }
}

func main() {
    r := gin.Default()

    r.GET("/ws", wsHandler)

    if err := r.Run(":8080"); err != nil {
        log.Fatal("Server failed to start:", err)
    }
}
