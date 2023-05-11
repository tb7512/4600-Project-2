package builtins

import (
	"strings"
	"fmt"
	"io"
)

func Echo(w io.Writer, args ...string) error { 
	userInput := args							//sets a variable to hold the input array
	fmt.Println(strings.Join(userInput, " "))	//Combines all of the array variables into a single string and prints
	return nil
}