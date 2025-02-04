package p238_product_of_array_except_self

// https://leetcode.com/problems/product-of-array-except-self/?envType=study-plan-v2&envId=top-interview-150
// 1,2,3,4
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = nums[i-1] * result[i-1]
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]

	}

	return result
}

func productExceptSelfOld(nums []int) []int {
	left := make([]int, len(nums), len(nums))
	right := make([]int, len(nums), len(nums))

	// left
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}
		prevVal := left[i-1]
		if i == 1 {
			left[i] = nums[i-1]
		} else {
			left[i] = prevVal * nums[i-1]
		}
	}

	// right
	for i := len(nums) - 1; i >= 0; i-- {
		if i+1 == len(nums) {
			continue
		}

		prevVal := right[i+1]
		if i == len(nums)-2 {
			right[i] = nums[i+1]
			left[i] += right[i]
		} else {
			right[i] = prevVal * nums[i+1]
			left[i] *= right[i]
		}
		if i == 0 {
			left[i] = right[i]
		}
	}

	return left
}
