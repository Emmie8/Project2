package builtins

import (
	"fmt"
	"os"
	//"path"
)

func PrintWorkingDirectory(args ...string) error {
	wd, err := os.Getwd() //wd stores the working directory path

	//if error doesn't equal 0(nil) then error was found and can't pwd
	if err != nil { 
		return fmt.Errorf("ERROR can't print working directory.")
	} else {
		fmt.Printf(wd) //prints working directory
	}

	return nil
}