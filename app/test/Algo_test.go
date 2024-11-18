package test

import "testing"

func TestAlgo1(t *testing.T) {
	productExceptSelf([]int{1, 2, 3, 4})

}
func productExceptSelf(nums []int) []int {
	leftMul := make([]int, len(nums))
	rightMul := make([]int, len(nums))

	for i, _ := range nums {
		if i == 0 {
			leftMul[0] = 1
		} else if i == len(nums)-1 {
			rightMul[i] = 1
			leftMul[i] = leftMul[i-1] * nums[i-1]
		} else if i >= 1 {
			leftMul[i] = leftMul[i-1] * nums[i-1]

		}

	}
	rightMul[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		rightMul[i] = rightMul[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = leftMul[i] * rightMul[i]
	}
	return nums
}
