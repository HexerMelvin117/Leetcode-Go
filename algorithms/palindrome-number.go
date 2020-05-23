package algorithms

import (
	"strconv"
	"strings"
)

func CheckPalindromeNumber(x int) bool {
	stringChars := strings.Split(strconv.Itoa(x), "")
	var reversedChars []string

	for i := len(stringChars) - 1; i >= 0; i-- {
		reversedChars = append(reversedChars, stringChars[i])
	}

	resultant := strings.Join(reversedChars, "")
	compareStr := strconv.Itoa(x)

	if resultant == compareStr {
		return true
	}
	return false
}
