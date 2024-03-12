package main

// solution with three main array
func productExceptSelf(nums []int) []int {
	var prefix, postfix, res []int
	mult := 1
	// prefix filling
	for i := range nums {
		mult *= nums[i]
		prefix = append(prefix, mult)
	}
	mult = 1
	// postfix filling
	for j := len(nums) - 1; j >= 0; j-- {
		mult *= nums[j]
		postfix = append(postfix, mult)
	}
	// postfix reverse
	i, j := 0, len(nums)-1
	for i = 0; i < len(nums)/2; i++ {
		postfix[i], postfix[j] = postfix[j], postfix[i]
		j--
	}
	for i := range nums {
		if i == 0 {
			res = append(res, postfix[i+1])
		} else if i == len(nums)-1 {
			res = append(res, prefix[i-1])
		} else {
			res = append(res, prefix[i-1]*postfix[i+1])
		}
	}
	return res
}
