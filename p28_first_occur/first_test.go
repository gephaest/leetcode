package p28_firstoccur

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	haystack := "abc"
	needle := "c"
	res := strStr(haystack, needle)

	fmt.Printf("res: %d\n", res)
}
