package p189_rotate_array

/*
Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
*/
func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		buffer := nums[0]

		for i := 0; i < len(nums); i++ {
			if i == 0 {
				nums[i] = nums[len(nums)-1]
			} else {
				nums[i], buffer = buffer, nums[i]
			}
		}
	}
}
func rotateSlow(nums []int, k int) {
	tmp := nums
	for i := 0; i < k; i++ {
		start := tmp[len(tmp)-1:]
		end := tmp[:len(tmp)-1]
		tmp = append(start, end...)
	}

	copy(nums, tmp)
}
