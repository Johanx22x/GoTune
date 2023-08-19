package api

import (
	"bufio"
	"database/sql"
	"fmt"
	"net"
	"strings"
)

// HandleConnection handles a TCP connection
func HandleConnection(conn net.Conn, db *sql.DB) {
    defer conn.Close()

    // Read incoming message, until EOF
    message, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        fmt.Println("Error reading message:", err)
        conn.Write([]byte("ERROR"))
        return
    }

    // Trim message
    message = strings.TrimSuffix(message, "\n")

    // Print message
    fmt.Println("Message received:", message)
    
    // Process message
    response, err := ProcessTCPRequest(message, db)
    if err != nil {
        fmt.Println("Error processing message:", err)
        conn.Write([]byte("ERROR"))
        return
    }

    // Write response
    conn.Write([]byte(response))
}
