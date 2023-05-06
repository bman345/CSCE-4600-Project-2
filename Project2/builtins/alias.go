package builtins

import (
	"fmt"
)

var aliases = make(map[string]string)

// Alias sets or prints aliases.
// If called with no arguments, it prints the current list of aliases.
// If called with arguments, it sets an alias.
func Alias(args ...string) error {
	if len(args) == 0 {
		// Print current list of aliases
		for alias, cmd := range aliases {
			fmt.Printf("%s='%s'\n", alias, cmd)
		}
		return nil
	}

	if len(args) == 1 {
		// Print the value of the given alias
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
