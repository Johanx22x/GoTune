package main

import (
	"errors"
	"fmt"
	"net"

	"github.com/Johanx22x/GoTune"
	"github.com/Johanx22x/GoTune/internal/api"
	"github.com/Johanx22x/GoTune/internal/db"
)

func main() {
    // Load config
    config := gotune.LoadConfig(true)

    // Connect to database
    db, err := repositories.NewDatabase(config.DatabaseFile)
    if err != nil {
        panic(errors.New("could not connect to database"))
    }
    fmt.Println("Database connection established")
    defer db.Close()

    // Start TCP server
    listener, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Port))
    if err != nil {
        panic(errors.New("could not start TCP server"))
    }
    fmt.Println("TCP server started")
    defer listener.Close()

    // Receive connections
    for {
        conn, err := listener.Accept()

        if err != nil {
            fmt.Println("could not accept connection")
            continue
        }

        // Deploy goroutine to handle connection
        go api.HandleConnection(conn, db)
    }
}
