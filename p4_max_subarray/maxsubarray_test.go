package p4_maxsubarray

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(arr)

	fmt.Printf("\nres: %d\n", res)
}
