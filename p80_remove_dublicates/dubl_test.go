package p80_remove_dublicates

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	//input := []int{1, 1, 1, 1, 2, 3, 3}
	input := []int{0, 0, 1, 1, 1, 1, 1, 2, 3, 3}

	fmt.Printf("in: %v\n", input)
	res := removeDuplicates(input)

	fmt.Printf("k: %d\n"+
		"res: %v", res, input[:res])
}
