package p14_longest_common_prefix

import "strings"

func maxLen(strs []string) int {
	maxL := 0
	for _, str := range strs {
		if len(str) >= maxL {
			maxL = len(str)
		}
	}

	return maxL
}

func longestCommonPrefix(strs []string) string {
	maxL := maxLen(strs)
	prefix := make([]string, maxL, maxL)

	for i := 0; i < maxL; i++ {
		for j, str := range strs {
			symbol := " "
			if len(str) > i {
				symbol = string(str[i])
			}

			if j == 0 {
				prefix[i] = string(symbol)
				//continue
			}

			if prefix[i] != string(symbol) {
				if i == 0 {
					return ""
				}
				return strings.Join(prefix[:i], "")
			}

		}
	}

	return strings.Join(prefix, "")
}
