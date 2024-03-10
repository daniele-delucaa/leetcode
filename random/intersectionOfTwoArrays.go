package main

// easy solution
func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	occ := make(map[int]bool)
	for i := range nums1 {
		occ[nums1[i]] = true
	}
	for i := range nums2 {
		if _, ok := occ[nums2[i]]; ok {
			res = append(res, nums2[i])
			delete(occ, nums2[i])
		}
	}
	return res
}

// easy solution but with array for the occurrences instead of the map
func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	occ := make([]int, 1000)
	for i := range nums1 {
		occ[nums1[i]]++
	}
	for i := range nums2 {
		if occ[nums2[i]] > 0 {
			res = append(res, nums2[i])
			occ[nums2[i]] = 0
		}
	}
	return res
}
