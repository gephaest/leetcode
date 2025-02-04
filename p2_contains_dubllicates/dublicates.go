package p2_contains_dubllicates

func containsDuplicate(nums []int) bool {
	valMap := make(map[int]bool)

	for _, current := range nums {
		if _, ok := valMap[current]; ok {
			return true
		}

		valMap[current] = true
	}

	return false
}
