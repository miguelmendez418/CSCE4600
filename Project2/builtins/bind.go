package builtins

import (
	"fmt"
	
)

//function simulates the bind command
func BindCommand(args ...string) error {
	if len(args) < 2 {
		return ErrInvalidArgCount
	}

	key := args[0]
	command := args[1]

	
	AliasMap[key] = command

	fmt.Printf("Binding '%s' to command '%s'\n", key, command)

	return nil
}
