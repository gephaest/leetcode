package p6_merge_strings_alternately

import (
	"fmt"
	"testing"
)

/*
Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r
*/

func TestName(t *testing.T) {
	res := mergeAlternately("ab", "pqrs")

	fmt.Printf("\n res: %s \n", res)
}
