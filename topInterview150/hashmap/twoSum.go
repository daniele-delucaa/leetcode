func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if index, ok := m[complement]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}