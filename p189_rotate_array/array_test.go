package p189_rotate_array

import (
	"fmt"
	"testing"
)

/*
Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
*/
func TestName(t *testing.T) {
	res := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(".", res)
	rotate(res, 3)
	fmt.Println(".", res)

}
