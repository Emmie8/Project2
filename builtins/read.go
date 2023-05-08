package builtins

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Read(args ...string) error {
	var userInput string
	var tempString string
	var err error

	if len(args) == 1 { // when a user enters a single variable
		if args[0] == "-d" { // checks if user entered an option without all parameters
			return fmt.Errorf("%w: -d requires an argument", ErrInvalidArgCount) // returns error and lets user enter command again
		}

		reader := bufio.NewReader(os.Stdin)                // gets full user input
		tempString, _ = reader.ReadString('\n')            // reads until newline
		userInput = string(tempString[:len(tempString)-1]) /// cuts newline from final string

		return os.Setenv(args[0], userInput) // adds the user input to the specified enviroment variable

	} else if len(args) >= 2 { // for multiple variables or if the user adds an option
		if args[0] == "-d" { // section that runs the code for the -d option
			var delim = []byte(args[1]) // gets the delimiter character as a byte

			reader := bufio.NewReader(os.Stdin)                // gets full user input
			tempString, _ = reader.ReadString(delim[0])        // stops at user specified delim character
			userInput = string(tempString[:len(tempString)-1]) // removes delim character from string
			if len(args) == 2 {                                // adds to default variable if none was specified
				return os.Setenv("REPLY", userInput)
			}

			newVariables := strings.SplitAfterN(userInput, " ", len(args)-2) // splits the user input string to match the number of user provided variables
			for i := 2; i < len(args); i++ {                                 // adds the new split strings to the enviroment variables
				err = os.Setenv(args[i], newVariables[i-2])
			}
			return err

		}

		//this section runs if the user only provides variables and no options
		reader := bufio.NewReader(os.Stdin)                // gets user input
		tempString, _ = reader.ReadString('\n')            // stops at newline
		userInput = string(tempString[:len(tempString)-1]) // gets rid of newline

		newVariables := strings.SplitAfterN(userInput, " ", len(args)) // splits full string into same number of strings as variables given
		for i := 0; i < len(args); i++ {                               // adds the strings to the enviroment variables specified
			err = os.Setenv(args[i], newVariables[i])
		}
		return err
	} else { // if the user just enters read
		reader := bufio.NewReader(os.Stdin)                // gets full input
		tempString, _ = reader.ReadString('\n')            // stops at newline
		userInput = string(tempString[:len(tempString)-1]) // gets rid of newline
		return os.Setenv("REPLY", userInput)               // adds user input to default enviroment variable
	}
}
