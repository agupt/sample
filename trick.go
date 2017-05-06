package sample

func twoSum(nums []int, target int) []int {
	diffs := make(map[int][]int)
	for i, num := range nums {
		if diff_indexs, ok := diffs[num]; ok {
			//number is reminder of previous number
			return []int{diff_indexs[0], i}
		}
		rem := target - num
		if _, ok := diffs[rem]; !ok {
			diffs[rem] = make([]int, 0)
		}
		diffs[rem] = append(diffs[rem], i)

	}
	return []int{}
}
