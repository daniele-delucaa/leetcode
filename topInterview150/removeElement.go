func removeElement(nums []int, val int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			res++
		} else {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != val {
					nums[i], nums[j] = nums[j], nums[i]
					res++
					break
				}
			}
		}
	}
	return res
}