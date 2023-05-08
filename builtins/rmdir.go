package builtins

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func IsEmpty(name string) bool {

	//open directory, and check if it opened properly
	//if not close it
	file, error := os.Open(name)
	if error != nil {
		return false
	}
	defer file.Close()

	//read names of files in directory
	_, error = file.Readdirnames(1)

	//check if it returned an eof error
	if error == io.EOF {
		return true
	}
	return false
}

func RemoveDirectory(args ...string) error {
	//Get the path from args[0]
	if len(args) == 0 {
		return fmt.Errorf("expected at least one arguments (directory)")
	}

	//get variable to decide number of loops
	loopOption := 0

	//current arg for path
	pathNum := 0

	//If first argument is -p, set path to next argument
	if args[0] == "-p" {
		loopOption = 1

		if len(args) < 2 {
			return fmt.Errorf("expected at least two arguments (directory)")
		}
		pathNum = 1
	}

	//if first arg is -p, loop one less time than length of args
	for i := 0; i < (len(args) - loopOption); i++ {

		//set path as string of arg
		path := args[pathNum]

		//Remove quotations
		if path[0] == '"' {
			path = path[1:]
		}
		if path[len(path)-1] == '"' {
			path = path[:len(path)-1]
		}

		//check if directory is empty
		empty := IsEmpty(path)

		if empty == false {
			return fmt.Errorf("directory must be empty")
		}

		//check if directory exists
		stat, err := os.Stat(path)

		if err != nil || !stat.IsDir() {
			return fmt.Errorf("directory does not exist")
		}

		//trying to delete directory
		if args[0] == "-p" {
			//get parent directory
			parent := filepath.Dir(path)

			//delete directory then parent
			err := os.Remove(path)
			if err != nil {
				return fmt.Errorf("failed to delete directory: %v", err)
			}

			err = os.Remove(parent)
			if err != nil {
				return fmt.Errorf("failed to delete directory: %v", err)
			}

		} else {
			//delete directory
			err := os.Remove(path)

			if err != nil {
				return fmt.Errorf("failed to remove directory: %v", err)
			}

		}

		fmt.Printf("Directory '%s' removed successfully.\n", path)

		//add one to pathNum
		pathNum += 1

	}
	return nil

}
