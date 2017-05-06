package sample

func lengthOfLongestSubstring(s string) int {
	l, _ := longestSubString(s)
	return l
}

func longestSubString(s string) (int, int) {
	var maxLen, index, cur_maxLen, cur_index int = 0, 0, 0, 0
	runeMap := make(map[rune]bool)
	for i, c := range s {
		if _, ok := runeMap[c]; ok {
			if maxLen < cur_maxLen {
				// adjust last know max lenght
				maxLen = cur_maxLen
				index = cur_index
			}
			// reset current
			cur_maxLen = 1
			cur_index = i
			runeMap = make(map[rune]bool)
			runeMap[c] = true
			continue
		}
		cur_maxLen++
		runeMap[c] = true
	}
	if maxLen < cur_maxLen {
		maxLen = cur_maxLen
		index = cur_index
	}
	return maxLen, index
}

func printLongest(s string, l, startIndex int) string {
	return s[startIndex : startIndex+l]
}
