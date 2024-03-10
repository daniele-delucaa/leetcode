package main

// easy solution
func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	occ := make(map[int]int)
	for i := range nums1 {
		occ[nums1[i]]++
	}
	for i := range nums2 {
		if _, ok := occ[nums2[i]]; ok {
			res = append(res, nums2[i])
			occ[nums2[i]]--
			if occ[nums2[i]] == 0 {
				delete(occ, nums2[i])
			}
		}
	}
	return res
}
