package main

func rotate(nums []int, k int) {
	k = k % len(nums) // we need to do this reassignment to handle the cases where k is greater or equal than the length of the array
	// if k is equal than the length of the array than the array doesn't need to rotate k times (the array remain the same after the rotations)
	// if k is greater than the length, we do the module, we are ensuring that k is within range [0, n)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(arr []int) {
	j := len(arr) - 1
	for i := 0; i < j; i++ {
		arr[i], arr[j] = arr[j], arr[i]
		j--
	}
}
