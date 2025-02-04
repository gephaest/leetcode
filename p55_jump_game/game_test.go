package p55_jump_game

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	nums := []int{0, 2, 3}

	res := canJump(nums)

	fmt.Printf("n: %v\n"+
		"ca: %v\n", nums, res)
}
