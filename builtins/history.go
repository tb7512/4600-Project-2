package builtins

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var history []string

// AddCommandToHistory adds a command to the history slice
func AddCommandToHistory(command string) {
	history = append(history, command)
}

// History prints out the command history
func History(w io.Writer, args ...string) error {
	// Print out the history slice
	for i, cmd := range history {
		fmt.Fprintf(w, "%d %s\n", i+1, cmd)
	}
	return nil
}
