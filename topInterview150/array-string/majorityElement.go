package main

func majorityElement(nums []int) int {
	occ := make(map[int]int)
	q := len(nums) / 2
	for i := range nums {
		occ[nums[i]]++
		if occ[nums[i]] > q {
			return nums[i]
		}
	}
	return 0
}

// soluzione con algoritmo di boyer moore
func majorityElement(nums []int) int {
	var res, count int
	for i := range nums {
		if count == 0 {
			res = nums[i]
		}
		if nums[i] == res {
			count++
		} else {
			count--
		}
	}
	return res
}
