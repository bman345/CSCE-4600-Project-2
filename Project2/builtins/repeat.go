package builtins

import (
	"fmt"
	"strconv"
	"os"
	"os/exec"
	"io"
)

// Function that checks if the given command to repeat is a valid command.
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

// The main function for the repeat command.
func Repeat(w io.Writer, args ...string) error {
    num, err := strconv.Atoi(args[0])
    if err != nil {
        return fmt.Errorf("%w: Input for loop number must be a positive number.", ErrInvalidArgCount)
    }
	
	// New name and args that are specifically for the looping command.
    nameR, argsR := args[1], args[2:]
	
	// Make sure that the given input is valid and if it is, repeat the given command "num" times.
    if num <= 0 {
        return fmt.Errorf("%w: Input for loop number cannot be zero or less than zero.", ErrInvalidArgCount)
    } else if num > 0 {
        for num > 0 {
            err := Check(w, nameR, argsR...) // Assign the error value returned by Check to a variable.
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

// Funciton specifically for running the command that is being repeated.
func executeCommand(nameR string, argsR ...string) error {
    cmd := exec.Command(nameR, argsR...)
    cmd.Stderr = os.Stderr
    return cmd.Run() // return the error value directly
}
