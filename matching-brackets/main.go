package main

import ( 
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	stringArg, readErr := reader.ReadString('\n')

	if readErr != nil {
		fmt.Printf("Error reading input: %s", readErr.Error())
		return
	}

	stringSlice := strings.Split(strigArg, "")
	isValid := validateBrackets(stringArg)

}

func validateBrackets (stringSlice []string) (bool, error) {
	
	openBracketIndex = -1

	for i, stringChar := range stringSlice {
		
		if isClosedBracket(stringChar) {
			if openBracketIndex == -1 {
				return false, errors.New(fmt.Sprintf('Unexpected %s at position %s'), stringChar, i)
			} else {
				//make sure open and close match type
			}

		} else if isOpenBracket(stringChar) {
			openBracketIndex = i
			isValid, validateErr := validateBrackets()

			if validateErr != nil {
				return false, validateErr
			}
		}
	}

	if openBracketIndex > -1 {
		return false, errors.New("Missing closing bracket")
	}

	return true, nil

}

func isClosedBracket(charString string) bool {

}

func isOpenBracket(charString string) bool {

}

//check if bracket/paren/brace
//if opening then find closing
//