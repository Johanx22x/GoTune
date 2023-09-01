package command

import (
	"errors"
	"fmt"
)

// Exit exits the program
func Exit() error {
    fmt.Println("Shutting down...")

    return errors.New("exit")
}
