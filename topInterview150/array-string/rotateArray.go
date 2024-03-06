package main

// non funziona, da rifare
func rotate(nums []int, k int) {
	var vals []int
	for i := range nums {
		if i >= len(nums)-k {
			vals = append(vals, nums[i])
		}
	}
	resIndex := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i < len(nums)-k {
			nums[resIndex] = nums[i]
			resIndex--
		}
	}
	for i := range vals {
		if i < k {
			nums[i] = vals[i]
		}
	}
}
