package p80_remove_dublicates

/*
Input: nums = [1,1,1,2,2,3]
Output: 5, nums = [1,1,2,2,3,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:

Input: nums = [1,1,1,1,2,3,3]
Output: 7, nums = [0,0,1,1,2,3,3,_,_]
Explanation: Your function should return k = 7, with the first seven elements of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
*/
func removeDuplicates1(nums []int) int {
	current := nums[0]
	counter := 1

	for i := 1; i < len(nums); i++ {
		if current == nums[i] {
			counter++
		} else { // next order, check old
			if counter > 2 {
				startRemoveFromIdx := i - 2
				nums = append(nums[startRemoveFromIdx:], nums[i:]...)
			} else {
				// skip
			}

		}

	}
	return 0
}

/*
1, 1, 1, 1, 2, 3, 3
[1], [1, 1, 2, 3, 3]
*/
func removeDuplicates(nums []int) int {
	finalLen := len(nums)
	currentVal := nums[0]
	counter := 1

	for i := 1; i < len(nums); i++ {
		value := nums[i]

		if currentVal == value {
			counter++
			if counter > 2 {
				nums = append(nums[:i-1], nums[i:]...)
				finalLen--
				i--
				continue
			}
		}

		if currentVal != value {
			currentVal = value
			counter = 1
			continue
		}

	}

	return finalLen
}
