package p3_longsubstr

func lengthOfLongestSubstring(s string) int {
	seenSet := map[string]struct{}{}
	maxLenght := 0

	if len(s) == 1 {
		return 1
	}

	subRune := []rune(s)
	left := 0

	for right, symbol := range subRune {
		symbolStr := string(symbol)

		if _, ok := seenSet[symbolStr]; ok {
			for {
				if _, ok := seenSet[symbolStr]; !ok {
					break
				}
				leftStrToRemove := string(subRune[left])
				delete(seenSet, leftStrToRemove)
				left++
			}
			seenSet[symbolStr] = struct{}{}
		} else {
			seenSet[symbolStr] = struct{}{}
			newMax := right - left + 1
			if newMax > maxLenght {
				maxLenght = newMax
			}
		}
	}

	return maxLenght
}

func lengthOfLongestSubstringOLD(s string) int {
	subRune := []rune(s)

	seen := map[string]int{}

	if len(s) == 1 {
		return 1
	}
inner:
	for idx, _ := range subRune {
		currentSeen := map[string]int{}
		for j := idx + 1; j <= len(subRune); j++ {
			subStr := s[idx:j]
			lastSymbol := subStr[len(subStr)-1:]

			// check if last symbol appeared
			if _, ok := currentSeen[lastSymbol]; ok {
				continue inner
			}

			currentSeen[lastSymbol] = 1

			if _, ok := seen[subStr]; ok {
				seen[subStr]++
			} else {
				seen[subStr] = 1
			}

		}
	}

	resStr := ""
	for subStr, _ := range seen {
		if len(subStr) >= len(resStr) {
			resStr = subStr
		}
	}

	return len(resStr)
}
