package p28_first_occur

import "strings"

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}

func strStrLoop(haystack string, needle string) int {
	if len(haystack) == 0 {
		return -1
	}

	if len(needle) == 0 {
		return -1
	}

	if haystack == needle {
		return 0
	}

	maxLoopSize := len(haystack) - len(needle) + 1

haystackLoop:
	for i := 0; i < maxLoopSize; i++ {
		rightSymbols := 0

		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				continue haystackLoop
			}

			rightSymbols++
		}

		if rightSymbols == len(needle) {
			return i
		}
	}

	return -1
}

func strStrEas(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
