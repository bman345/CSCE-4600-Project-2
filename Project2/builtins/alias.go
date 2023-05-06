package builtins

import (
	"fmt"
)

var aliases = make(map[string]string)


func Alias(args ...string) error {
	if len(args) == 0 {
		// This will print current list of aliases
		for alias, cmd := range aliases {
			fmt.Printf("%s='%s'\n", alias, cmd)
		}
		return nil
	}

	if len(args) == 1 {
		// this will print the value of the given alias
		cmd, ok := aliases[args[0]]
		if !ok {
			return fmt.Errorf("alias: %s not found", args[0])
		}
		fmt.Printf("%s='%s'\n", args[0], cmd)
		return nil
	}

	// Set an alias
	aliases[args[0]] = args[1]
	return nil
}
