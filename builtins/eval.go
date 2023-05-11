package builtins


import (
	"io"
	"os/exec"
	"strings"
)

// eval reads the specified arguments as input to the shell and executes the resulting command or commands.
func eval(w io.Writer, args ...string) error {
	// Concatenate the input arguments into a single string.
	input := strings.Join(args, " ")

	// Split the input separate the command name and the command arguments.
	cmds := strings.Split(input, "|")

	// Execute each command in the pipeline.
	var (
		stdin  io.Reader = nil
		stdout io.Writer = w
		err    error     = nil
	)
	for _, cmd := range cmds {
		// Trim any leading/trailing white space.
		cmd = strings.TrimSpace(cmd)

		// Split the command into its parts.
		parts := strings.Fields(cmd)
		if len(parts) == 0 {
			continue // Skip empty commands.
		}

		// Create the command object.
		cmdObj := exec.Command(parts[0], parts[1:]...)
		cmdObj.Stdin = stdin
		cmdObj.Stdout = stdout

		// Execute the command and wait for it to complete.
		if err = cmdObj.Start(); err != nil {
			return err
		}
		if stdin != nil {
			if closer, ok := stdin.(io.Closer); ok {
				closer.Close()
			}
		}
		if stdin, err = cmdObj.StdoutPipe(); err != nil {
			return err
		}
		if stdout != nil {
			if closer, ok := stdout.(io.Closer); ok {
				closer.Close()
			}
		}
		if stdout, err = cmdObj.StdinPipe(); err != nil {
			return err
		}
		if err = cmdObj.Wait(); err != nil {
			return err
		}
	}

	return nil
}
