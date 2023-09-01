package command

import (
    "fmt"
)


func Help() error {
    fmt.Println("help - Display this help message")
    fmt.Println("exit - Exit the program")
    fmt.Println("add - Add a new task")

    return nil
}
