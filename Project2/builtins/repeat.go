package builtins

import (
	"fmt"
	"strconv"
	"os"
	"os/exec"
	"io"
	
	//"github.com/jh125486/CSCE4600/Project2/builtins"
)

func Check(w io.Writer, nameR string, argsR ...string) error {
    switch nameR {
    case "cd":
        return ChangeDirectory(argsR...)
    case "env":
        return EnvironmentVariables(w, argsR...)
    case "echo":
        return Echo(argsR...)
    case "alias":
        return Alias(argsR...)
    }
    return executeCommand(nameR, argsR...)
}

func Repeat(w io.Writer, args ...string) error {
    num, err := strconv.Atoi(args[0])
    if err != nil {
        return fmt.Errorf("%w: Input for loop number must be a positive number.", ErrInvalidArgCount)
    }

    nameR, argsR := args[1], args[2:]

    if num <= 0 {
        return fmt.Errorf("%w: Input for loop number cannot be zero or less than zero.", ErrInvalidArgCount)
    } else if num > 0 {
        for num > 0 {
            err := Check(w, nameR, argsR...) // assign the error value returned by Check to a variable
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
            num--
        }
    } else {
        return fmt.Errorf("%w: Input for loop number must be a positive number.", ErrInvalidArgCount)
    }
    return nil
}

func executeCommand(nameR string, argsR ...string) error {
    cmd := exec.Command(nameR, argsR...)
    cmd.Stderr = os.Stderr
    return cmd.Run() // return the error value directly
}
