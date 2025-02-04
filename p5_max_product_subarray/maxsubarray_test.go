package _maxsubarray

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	arr := []int{-2}
	res := maxProduct(arr)

	fmt.Printf("\nres: %d\n", res)
}
