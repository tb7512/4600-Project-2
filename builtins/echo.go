package builtins

import (
	"strings"
	"fmt"
	"io"
)


func Echo(w io.Writer, args ...string) error {
	userInput := args							//sets a variable to hold the input array
	for i := 0; i < len(args)/2 - 1; i++ {
		userInput = append(userInput, args[i])
	}
	fmt.Println(strings.Join(userInput, " "))	//Combines all of the array variables into a single string
	return nil
}