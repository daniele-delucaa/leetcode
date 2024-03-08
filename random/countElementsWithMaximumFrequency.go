func maxFrequencyElements(nums []int) int {
	var res int
	occ := make(map[int]int)
	for i := range nums {
		occ[nums[i]]++
	}
	max := 0
	for _, v := range occ {
		if v > max {
			max = v
		}
	}
	for _, v := range occ {
		if v == max {
			for i := 0; i < v; i++ {
				res++
			}
		}
	}
	return res
}