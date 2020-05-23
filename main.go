package main

import (
	"fmt"

	"github.com/hexermelvin117/goalgorithms/algorithms"
)

func main() {
	myNums := []int{2, 11, 15}
	fmt.Println("WE Good")
	myResult := algorithms.TwoSum(myNums, 26)
	fmt.Println(myResult)

	reversedResult := algorithms.ReverseInteger(455234)
	fmt.Println(reversedResult)
}
