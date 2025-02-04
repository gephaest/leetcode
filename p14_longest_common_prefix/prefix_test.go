package p14_longest_common_prefix

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	in := []string{"flower", "flower", "flower", "flower"}

	res := longestCommonPrefix(in)

	fmt.Printf("in: %v\n"+
		"out: %s\n", in, res)
}
