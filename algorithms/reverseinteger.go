package algorithms

import (
	"fmt"
	"strconv"
	"strings"
)

func ReverseInteger(x int) int {
	stringChars := strings.Split(strconv.Itoa(x), "")
	var reversedChars []string

	if stringChars[0] == "-" {
		reversedChars = append(reversedChars, stringChars[0])
		for i := len(stringChars) - 1; i >= 1; i-- {
			reversedChars = append(reversedChars, stringChars[i])
		}
	} else {
		for i := len(stringChars) - 1; i >= 0; i-- {
			reversedChars = append(reversedChars, stringChars[i])
		}
	}

	resultantString := strings.Join(reversedChars, "")
	reversedInt, err := strconv.Atoi(resultantString)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	// Checks if it is an int32
	if reversedInt >= 2147483647 || reversedInt <= -2147483647 {
		return 0
	}

	return reversedInt
}
