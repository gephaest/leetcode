package p238_product_of_array_except_self

import (
	"fmt"
	"testing"
)

// 0,0,9,0,0
func TestName(t *testing.T) {
	nums := []int{-1, 1, 0, -3, 3±}
	//nums := []int{1, 2, 3, 4}
	//nums := []int{3, 4, 5, 6}

	res := productExceptSelf(nums)

	fmt.Printf("nums: %v\n"+
		"res: %v\n", nums, res)
}
