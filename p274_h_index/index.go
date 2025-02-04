package p274_h_index

import (
	"sort"
)

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	result := 0
	if len(citations) == 0 {
		return 0
	}

	for idx, citationsForPaper := range citations {
		if citationsForPaper >= idx+1 {
			result++
		} else {
			break
		}

	}

	return result
}

//citationsPerspective := len(citations) - idx + 1
//
//if idx >= citationsPerspective {
//	break
//}
//result++
