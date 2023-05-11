package builtins

import (
	"fmt"
)

func Echo(args ...string) error {
	fmt.Printf("%s \n", args)
	return nil
}
