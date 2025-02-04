package p55_jump_game

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	finish := len(nums) - 1
	farPointIdx := 0

	for i := 0; i < len(nums)-1; i++ {
		if farPointIdx < i {
			return false
		}

		currentFarIdx := i + nums[i]
		// reach finish
		if currentFarIdx >= finish {
			return true
		}
		if currentFarIdx > farPointIdx {
			farPointIdx = currentFarIdx
		}
	}

	if farPointIdx < finish {
		return false
	}

	return false
}
