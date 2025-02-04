package p6_merge_strings_alternately

/*
Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r
*/
func mergeAlternately(word1 string, word2 string) string {
	w1len := len(word1)
	w1rune := []rune(word1)
	w2rune := []rune(word2)
	w2len := len(word2)
	maxLen := max(w1len, w2len)

	res := []rune("")

	for i := 0; i < maxLen; i++ {
		if i < w1len {
			res = append(res, w1rune[i])
		}
		if i < w2len {
			res = append(res, w2rune[i])
		}
	}
	return string(res)
}
