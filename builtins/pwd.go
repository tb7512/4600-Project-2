package builtins

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	ErrInvalidArg = errors.New("invalid argument")
)

func workingDirectory(w io.Writer, args ...string) error {
	if len(args) > 1 {
		return fmt.Errorf("Too many arguments")
	}
	if len(args) == 1 {
		if args[0] == "-P" {
			workingDirectory, err := os.Getwd()
			if err != nil {
				return err
			}
			workingPath, err := filepath.EvalSymlinks(workingDirectory)
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(w, workingPath)
			return err
		} else if args[0] != "-L" {
			return fmt.Errorf("pwd only accepts -L or -P as arguments")
		}
	}
	pwd := os.Getenv("PWD")
	_, err := fmt.Fprintln(w, pwd)
	return err
}
