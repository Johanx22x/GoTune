package main

import (
	"errors"
	"fmt"
	"log"
	"net"
    "sync"

	"github.com/Johanx22x/GoTune"
	"github.com/Johanx22x/GoTune/internal/api"
	"github.com/Johanx22x/GoTune/internal/command"
	"github.com/Johanx22x/GoTune/internal/db"
)

func main() {
    // Load config
    config := gotune.LoadConfig(true)

    // Connect to database
    db, err := repositories.NewDatabase(config.DatabaseFile)
    if err != nil {
        log.Fatal(errors.New("Could not connect to database due to: " + err.Error()))
    }
    defer db.Close()

    // Start TCP server
    listener, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Port))
    if err != nil {
        log.Fatal(errors.New("Could not start TCP server due to: " + err.Error()))
    }
    defer listener.Close()

    // Receive connections
    var wg sync.WaitGroup
    go func() {
        for {
            conn, err := listener.Accept()

            if err != nil {
                log.Println("Could not accept connection due to: ", err.Error())
                continue
            }

            wg.Add(1)

            // Deploy goroutine to handle connection
            go func() {
                defer wg.Done()
                api.HandleConnection(conn, db)
            }()
        }
    }()

    fmt.Println("GoTune v" + config.Version + " is now listening for connections on port " + config.Port)
    fmt.Println("Type 'help' for a list of commands")
    for command.Listen() { }

    // Wait for all goroutines to finish
    wg.Wait()
}
