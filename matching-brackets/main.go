package main

import ( 
	"bufio"
	"os"
	"fmt"
	"strings"
	"errors"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	stringArg, readErr := reader.ReadString('\n')

	if readErr != nil {
		fmt.Printf("Error reading input: %s", readErr.Error())
		return
	}

	stringSlice := strings.Split(stringArg, "")
	isValid, err := validateBrackets(stringSlice)

	if err != nil {
		fmt.Printf("Invalid Bracket usage: %s", err.Error())
	} else if isValid {
		fmt.Println("Valid Brackets!")
	}

}


func validateBrackets(stringSlice []string) (bool, error) {
	
	openBracketIndex := -1

	for i, stringChar := range stringSlice {
		
		if isClosed, bracketType := isClosedBracket(stringChar); isClosed {
			fmt.Printf("BracketType open: %s, openBracketIndex: %d\n", bracketType.Open, openBracketIndex)
			if (openBracketIndex > -1) && (bracketType.Open == stringSlice[openBracketIndex]) {
				fmt.Println("Closed bracket")
				openBracketIndex = -1
			} else {
				
				return false, errors.New(fmt.Sprintf("Unexpected %s ", stringChar, i))
			}
		} else if isOpen,_ := isOpenBracket(stringChar); isOpen {
			if openBracketIndex == -1 {
				openBracketIndex = i
			} else {
				_, validateErr := validateBrackets(stringSlice[i:])
	
				if validateErr != nil {
					return false, validateErr
				}
			}

		}
	}

	if openBracketIndex > -1 {
		return false, errors.New("Missing closing bracket")
	}

	return true, nil

}



func isClosedBracket(charString string) (bool, Bracket) {
	for _,bracket := range bracketTypes {
		if bracket.Close == charString {
			return true, bracket
		}
	}
	return false, Bracket{}
}

func isOpenBracket(charString string) (bool, Bracket) {
	for _,bracket := range bracketTypes {
		if bracket.Open == charString {
			return true, bracket
		}
	}
	return false, Bracket{}
}



type Bracket struct {
	Open string
	Close string
	Type string
}

var bracketTypes = []Bracket {
	Bracket {
		Open: "{",
		Close: "}",
		Type: "curly",
	},
	Bracket {
		Open: "[",
		Close: "]",
		Type: "square",
	},
	Bracket {
		Open: "(",
		Close: ")",
		Type: "paren",
	},
}