package sample

func twoSum(nums []int, target int) []int {
	diffs := make(map[int][]int)
	for i, num := range nums {
		if diffIndexs, ok := diffs[num]; ok {
			//number is reminder of previous number
			return []int{diffIndexs[0], i}
		}
		rem := target - num
		if _, ok := diffs[rem]; !ok {
			diffs[rem] = make([]int, 0)
		}
		diffs[rem] = append(diffs[rem], i)

	}
	return []int{}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	count := len(nums1) + len(nums2)
	odd := true
	middleIndex := count / 2
	if count%2 == 0 {
		odd = false
	}
	var i, j, index, previous, current int = 0, 0, 0, 0, 0
	for index < count {
		if i == len(nums1) && j < len(nums2) {
			// nums1 finished continue travesing nums2
			previous = current
			current = nums2[j]
			j++
		} else if j == len(nums2) && i < len(nums1) {
			// nums2 finished continue travesing nums1
			previous = current
			current = nums1[i]
			i++
		} else if i < len(nums1) && j < len(nums2) && nums1[i] <= nums2[j] {
			// nums1 is smaller use it
			previous = current
			current = nums1[i]
			i++
		} else if i < len(nums1) && j < len(nums2) && nums1[i] > nums2[j] {
			// nums2 is smaller use it
			previous = current
			current = nums2[j]
			j++
		} else {
			panic("")
		}
		// check if have reached mid value
		if odd && index == middleIndex {
			//fmt.Println("Current: ", current, " Index: ", index)
			return float64(current)
		}
		if !odd && index == middleIndex {
			//fmt.Println("Current: ", current, " Previous: ", previous, " Index: ", index)
			return float64((current + previous)) / 2
		}
		//fmt.Println("Index: ", index, "Current: ", current, " Previous: ", previous, " i: ", i, " j: ", j)
		index++
	}
	return 0
}
