package builtins

import (
	"os"
	"os/exec"
)

//function that simulates command
func ExecCommand(args ...string) error {
	if len(args) < 1 {
		return ErrInvalidArgCount
	}

	cmd := exec.Command(args[0], args[1:]...)

	//Set correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	//Execute command to return the error
	return cmd.Run()
}
