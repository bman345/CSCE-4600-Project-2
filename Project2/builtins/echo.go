package builtins

import (
	"fmt"
	"os"
)

func Echo(args...string){
  for _, userInput := range os.Args[1:]{
		fmt.Printf("%s", userInput)
	} 
}
