package p274_h_index

import (
	"fmt"
	"testing"
)

/*
Example 1:

Input: citations = [3,0,6,1,5]
Output: 3
Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had received 3, 0, 6, 1, 5 citations respectively.
Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.
Example 2:

Input: citations = [1,3,1]
Output: 1
*/
func TestName(t *testing.T) {
	//citations := []int{0}
	citations := []int{3, 0, 6, 1, 5}

	res := hIndex(citations)

	fmt.Printf("citations: %v\n"+
		"result: %d\n", citations, res)
}
