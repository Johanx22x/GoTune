package main

import (
	"errors"
	"fmt"

	"github.com/Johanx22x/GoTune"
	"github.com/Johanx22x/GoTune/internal/db/repositories"
)

func main() {
    config := gotune.LoadConfig(true)

    db, err := repositories.NewDatabase(config.DatabaseFile)
    if err != nil {
        panic(errors.New("could not connect to database"))
    }
    fmt.Println("Database connection established")
    defer db.Close()
}
