package builtins

import (
	"fmt"
	"os"
)

//to print current working directory
func PrintWorkingDirectory() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(wd)
}
