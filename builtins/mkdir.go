package builtins

import (
	"fmt"
	"os"
)

func CreateDirectory(args ...string) error {
	//Get the path from args[0]
	if len(args) == 0 {
		return fmt.Errorf("expected one or two arguments (directory)")
	}

	// Store the path as a string
	path := args[0]

	//If path == -p, set path to args[1]
	if path == "-p" {
		if len(args) != 2 {
			return fmt.Errorf("expected two arguments (directory)")
		}
		path = args[1]
	}

	//Remove quotations
	if path[0] == '"' {
		path = path[1:]
	}
	if path[len(path)-1] == '"' {
		path = path[:len(path)-1]
	}


	// Check if the directory already exists
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return fmt.Errorf("directory already exists")
	}

	// Create the directory
	if(args[0] == "-p"){
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
		
	}else{
	err := os.Mkdir(path, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
}

	fmt.Printf("Directory '%s' created successfully.\n", path)
	return nil
}