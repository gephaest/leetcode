package p13_roman_to_integer

import (
	"strings"
)

func romanToInt(s string) int {
	res := 0
	numbersMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && numbersMap[s[i]] < numbersMap[s[i+1]] {
			res -= numbersMap[s[i]]
		} else {
			res += numbersMap[s[i]]
		}
	}

	return res
}

func romanToIntOLD2(s string) int {
	numbersMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	for i, char := range s {
		if v, ok := numbersMap[char]; ok {
			result += v
		}

		if i == 0 {
			continue
		}

		prevRune := s[i-1]

		switch char {
		case 'V':
			if prevRune == 'I' {
				result += 4 - numbersMap['V'] - numbersMap['I']
				continue
			}

		case 'X':
			if prevRune == 'I' {
				result += 9 - numbersMap['X'] - numbersMap['I']
				continue
			}

		case 'L':
			if prevRune == 'X' {
				result += 40 - numbersMap['L'] - numbersMap['X']
				continue
			}
		case 'C':
			if prevRune == 'X' {
				result += 90 - numbersMap['C'] - numbersMap['X']
				continue
			}
		case 'D':
			if prevRune == 'C' {
				result += 400 - numbersMap['D'] - numbersMap['C']
				continue
			}
		case 'M':
			if prevRune == 'C' {
				result += 900 - numbersMap['M'] - numbersMap['C']
				continue
			}
		}
	}

	return result
}

func romanToIntOLD(s string) int {
	result := 0
	for i, char := range s {

		// 1
		if strings.ContainsRune("I", char) {
			result += 1
		}

		// 5
		if strings.ContainsRune("V", char) {
			if i > 0 && strings.ContainsRune("I", []rune(s)[i-1]) {
				result += 3 // IV = 1,3=4
			} else {
				result += 5
			}
		}

		// 10
		if strings.ContainsRune("X", char) {
			if i > 0 && strings.ContainsRune("I", []rune(s)[i-1]) {
				result += 8 //1,8=9
			} else {
				result += 10
			}
		}

		// 50
		if strings.ContainsRune("L", char) {
			if i > 0 && strings.ContainsRune("X", []rune(s)[i-1]) {
				result += 30 //10,30=40
			} else {
				result += 50
			}
		}

		// 100
		if strings.ContainsRune("C", char) {
			if i > 0 && strings.ContainsRune("X", []rune(s)[i-1]) {
				result += 80 // 10,80=90
			} else {
				result += 100
			}
		}

		// 500
		if strings.ContainsRune("D", char) {
			if i > 0 && strings.ContainsRune("C", []rune(s)[i-1]) {
				result += 300 // 100,400=500
			} else {
				result += 500
			}
		}

		// 1000
		if strings.ContainsRune("M", char) {
			if i > 0 && strings.ContainsRune("C", []rune(s)[i-1]) {
				result += 800 // 100,800=900
			} else {
				result += 1000
			}
		}
	}

	return result
}
