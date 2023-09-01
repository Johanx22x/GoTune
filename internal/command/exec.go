package command

import (
    "errors"
    "fmt"
)

func Exec(command string, args []string) error {
    return errors.New(fmt.Sprintf("command '%v' not found", command))
}
