package main

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums) - 1
	}
	l, r := 1, 2
	for i := 1; i < len(nums); i++ {
		if nums[l] == nums[r] {
			r++
		} else {
			l++
			nums[l+1] = nums[r+1]
			r++
		}
	}
	return l
}
