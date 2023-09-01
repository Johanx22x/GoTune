package command

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "log"
    "errors"
)

// Listen listens for input from the user
func Listen() bool {
    // Print the prompt
    fmt.Print(">>> ")

    // Create a new reader
    reader := bufio.NewReader(os.Stdin)

    // Read the input
    input, err := reader.ReadString('\n')
    if err != nil {
        log.Println(errors.New("could not read input due to: " + err.Error()))
        return true
    }

    // Trim the input
    input = strings.Trim(input, "\n")

    // Handle the input
    err = HandleInput(input)
    if err == nil {
        return true
    }

    switch err.Error() {
    case "exit":
        return false
    default:
        log.Println(errors.New("could not handle input due to: " + err.Error()))
        return true
    }
}
