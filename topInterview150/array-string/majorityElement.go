package main

func majorityElement(nums []int) int {
	var res int
	occ := make(map[int]int)
	q := len(nums) / 2
	for i := range nums {
		occ[nums[i]]++
		if occ[nums[i]] > q {
			res = nums[i]
		}
	}
	return res
}
