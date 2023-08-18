package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Johanx22x/GoTune"
	"github.com/Johanx22x/GoTune/internal/api"
	"github.com/Johanx22x/GoTune/internal/db"
	"github.com/gorilla/mux"
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

    // Create router
    router := mux.NewRouter()

    // Register routes
    api.RegisterRoutes(router, db)

    // Serve the API on designated port
    fmt.Println("Serving API on port " + config.Port)
    http.Handle("/", router)
    http.ListenAndServe(":"+config.Port, nil)
}
