package sample

import (
	"testing"
)

func Test_FindMedianSortedArrays(t *testing.T) {
	ok := func(nums1, nums2 []int, median float64) {
		m := findMedianSortedArrays(nums1, nums2)
		if m != median {
			t.Error("Expected median ", median, " got ", m)
		}
	}
	ok([]int{}, []int{}, 0)
	ok([]int{}, []int{1}, 1)
	ok([]int{1}, []int{}, 1)
	ok([]int{1}, []int{1}, 1)
	ok([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, 2.5)
	ok([]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, 4.5)
	ok([]int{5, 6, 7, 8}, []int{1, 2, 3, 4}, 4.5)
	ok([]int{1, 3, 5, 6}, []int{2, 4, 7, 8}, 4.5)
}
