package builtins

import (
	"fmt"
	"io"
	"strconv"
)

var history []string

// AddCommandToHistory adds a command to the history slice
func AddCommandToHistory(command string) {
	history = append(history, command)
}

// History prints out the command history
func History(w io.Writer, args ...string) error {
	// Parse the options and arguments.
	var (
		clear  bool
		printR bool
		printH bool
		n      int
		err    error
	)
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-c":
			clear = true
		case "-r":
			printR = true
		case "-h":
			printH = true
		default:
			// Parse the argument as an integer.
			if n, err = strconv.Atoi(args[i]); err != nil {
				return fmt.Errorf("invalid argument: %s", args[i])
			}
		}
	}

	// Handle the options.
	if clear {
		// Clear the history.
		history = make([]string, 0)
	} else if printH {
		// Print the history without event numbers.
		for _, cmd := range history {
			_, _ = fmt.Fprintln(w, cmd)
		}
	} else {
		// Print the history with event numbers.
		if printR {
			// Reverse the order of the history.
			for i := len(history) - 1; i >= 0; i-- {
				_, _ = fmt.Fprintf(w, "%d %s\n", i+1, history[i])
			}
		} else {
			for i, cmd := range history {
				_, _ = fmt.Fprintf(w, "%d %s\n", i+1, cmd)
			}
		}
	}

	// Display a specific number of commands.
	if n > 0 {
		if n > len(history) {
			n = len(history)
		}
		if printR {
			// Print the last n commands in reverse order.
			for i := len(history) - n; i < len(history); i++ {
				_, _ = fmt.Fprintf(w, "%d %s\n", i+1, history[i])
			}
		} else {
			for i := 0; i < n; i++ {
				_, _ = fmt.Fprintf(w, "%d %s\n", i+1, history[i])
			}
		}
	}

	return nil
}