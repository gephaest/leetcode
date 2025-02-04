package p4_max_subarray

import "math"

func maxSubArray(nums []int) int {
	currentSum := 0
	maxSum := math.MinInt32

	for _, value := range nums {
		if (currentSum + value) > value {
			currentSum = currentSum + value
		} else {
			currentSum = value
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
