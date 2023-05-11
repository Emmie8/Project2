// write code for ls builtin
package builtins

import (
	"fmt"
	"io/ioutil"
)

func List(args ...string) error {
	//if no arguments are passed in, then list all files in current directory
	if len(args) == 0 {
		files, err := ioutil.ReadDir(".")
		if err != nil {
			return fmt.Errorf("ERROR: can't list files in current directory")
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
		return nil
	}

	//if arguments are passed in, then list all files in the directory specified
	if len(args) == 1 {
		files, err := ioutil.ReadDir(args[0])
		if err != nil {
			return fmt.Errorf("ERROR: can't list files in directory")
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
		return nil
	}

	//if more than one argument is passed in, then list all files in the directories specified
	if len(args) > 1 {
		for _, arg := range args {
			files, err := ioutil.ReadDir(arg)
			if err != nil {
				return fmt.Errorf("ERROR: can't list files in directory")
			}
			for _, file := range files {
				fmt.Println(file.Name())
			}
		}
		return nil
	}

	return nil
}
