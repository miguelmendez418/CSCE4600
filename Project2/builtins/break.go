package builtins

import (
	"errors"
	
)


var ErrBreak = errors.New("break command")

//function simulates break command
func BreakCommand() error {
	return ErrBreak
}
