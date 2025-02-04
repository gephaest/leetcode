package p01_twosum

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	res := twoSum([]int{-1, -2, -3, -4, -5}, -8)
	fmt.Printf("\n%+v\n", res)
}
