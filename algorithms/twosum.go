package algorithms

func TwoSum(nums []int, target int) []int {
	var myResult []int
	for i := 0; i < len(nums); i++ {
		for k := 0; k < len(nums); k++ {
			if k == i {
				k = k + 1
			}
			resultSum := 0

			if k >= len(nums) {
				resultSum = 0
			} else {
				resultSum = nums[i] + nums[k]
			}
			if resultSum == target {
				myResult = append(myResult, i, k)
			}
		}
	}
	if len(myResult) > 0 {
		return myResult[0:2]
	}
	return myResult
}
