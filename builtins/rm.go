// Dajanek Davis
package builtins

import (
	"fmt"
	"os"
)

func RemoveFiles(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("not enough arguments")
	}

	fname := args[0]

	err := os.Remove(fname)
	if err != nil {
		return err
	}

	return nil

}
