package command

import (
    "strings"
)

// HandleInput takes the input from the user and calls the appropriate function
func HandleInput(input string) error {
    // Split the input into a slice
    splitInput := strings.Split(input, " ")

    // Get the command
    command := splitInput[0]

    switch command {
    case "help":
        return Help()
    case "exit":
        return Exit()
    default:
        return Exec(command, splitInput[1:])
    }
}
