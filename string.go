package sample

func lengthOfLongestSubstring(s string) int {
	_, _, l := longestSubString(s)
	return l
}

func longestSubString(s string) (int, int, int) {
	charMap := make(map[rune]int)
	var lStart, lEnd, lLen int // start & end index of current longgest substring
	var cStart, cEnd, cLen int // start & end index of current substring under investigation
	for i, r := range s {
		if chIndex, ok := charMap[r]; ok && chIndex >= cStart {
			// found dup after current start index; update and swtich to next substring
			if cLen > lLen {
				// substring under investigation is current longest
				lStart = cStart
				lEnd = cEnd
				lLen = cLen
			}
			cStart = chIndex + 1
		}
		charMap[r] = i
		cEnd = i
		cLen = cEnd - cStart + 1
	}
	if cLen > lLen {
		lStart = cStart
		lEnd = cEnd
		lLen = cLen
	}
	return lStart, lEnd, lLen
}

func printLongest(s string, l, startIndex int) string {
	return s[startIndex : startIndex+l]
}
