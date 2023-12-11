package builtins

import (
	"fmt"
)


var AliasMap = make(map[string]string)

//creates alias for a command
func SetAlias(alias, command string) {
	AliasMap[alias] = command
}

//print current aliases
func PrintAliases() {
	fmt.Println("Current Aliases:")
	for alias, command := range AliasMap {
		fmt.Printf("%s=%s\n", alias, command)
	}
}
