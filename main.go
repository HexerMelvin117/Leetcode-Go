package main

import (
	"fmt"

	"github.com/HexerMelvin117/Leetcode-Go/algorithms"
)

func main() {
	myNums := []int{2, 11, 15}
	fmt.Println("WE Good")

	// returns indexes of two numbers that equal to the second argument
	myResult := algorithms.TwoSum(myNums, 26)
	fmt.Println(myResult)

	// must be an int32
	reversedResult := algorithms.ReverseInteger(455234)
	fmt.Println(reversedResult)

	isPalindrome := algorithms.CheckPalindromeNumber(22)
	fmt.Println(isPalindrome)
}
