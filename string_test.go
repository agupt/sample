package sample

import (
	"testing"
)

func Test_LongestString(t *testing.T) {
	ok := func(str, expString string, expStart, expEnd, expLen int) {
		s, e, l := longestSubString(str)
		if expStart != s || expEnd != e || expLen != l {
			t.Fatalf("For String %s Expected start %d:%d end %d:%d len %d:%d", str, expStart, s, expEnd, e, expLen, l)
		}
		lString := printLongest(str, l, s)
		if lString != expString {
			t.Errorf("Expected %s got %s", expString, lString)
		}
	}
	ok("dvdm", "vdm", 1, 3, 3)
	ok("abcabcbb", "abc", 0, 2, 3)
	ok("bbbb", "b", 0, 0, 1)
	ok("pwwkew", "wke", 2, 4, 3)
	ok("abba", "ab", 0, 1, 2)
}
