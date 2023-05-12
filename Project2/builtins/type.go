package builtins

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Type(args ...string) error{
	
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "type: missing operand\n")
		return fmt.Errorf("missing operand")
	}
	name := args[0]
	path, err := LookPath(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "type: %v\n", err)
		return err
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "type: %v\n", err)
		return err
	}
	if fileInfo.IsDir() {
		fmt.Printf("%s is a directory\n", name)
	} else {
		fmt.Printf("%s is %s\n", name, path)
	}
	return nil
}

// Returns the path of the given executable file name
func LookPath(file string) (string, error) {
	_, err := os.Stat(file)
	if err == nil {
		return file, nil
	}
	pathEnv := os.Getenv("PATH")
	paths := strings.Split(pathEnv, string(os.PathListSeparator))
	for _, dir := range paths {
		path := filepath.Join(dir, file)
		_, err := os.Stat(path)
		if err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("%s: command not found", file)
}
