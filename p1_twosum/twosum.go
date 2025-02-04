package p01_twosum

// res := twoSum([]int{-3,4,3,90}, 0)
// [0,2]
func twoSum(nums []int, target int) []int {
	mMap := map[int]int{}

	for i, iValue := range nums {
		possibleAnswer := target - iValue
		if answerIdx, ok := mMap[possibleAnswer]; ok {
			return []int{answerIdx, i}
		}

		mMap[iValue] = i
	}
	return []int{}
}
