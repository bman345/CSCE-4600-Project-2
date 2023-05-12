package builtins

import (
	"fmt"
	"os"
)

func Pwd(args ...string) error {
	//variables for getting directory
	workingDir, error := os.Getwd()

	//checks if error occurs, if not prints directory
	if error != nil {
		fmt.Printf("%s \n", error)
	}

	fmt.Printf("%s \n", workingDir)

	return nil
}
