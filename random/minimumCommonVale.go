package main

import "math"

// easy solution with map
func getCommon(nums1 []int, nums2 []int) int {
	occ := make(map[int]bool)
	var min int = math.MaxInt
	for i := range nums1 {
		occ[nums1[i]] = true
	}
	for i := range nums2 {
		if _, ok := occ[nums2[i]]; ok {
			if nums2[i] < min {
				min = nums2[i]
			}
		}
	}
	if min == math.MaxInt {
		return -1
	}
	return min
}

// easy solution but doesnt work with very long array (time limit exceeded)
func getCommon(nums1 []int, nums2 []int) int {
	for i := range nums1 {
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				return nums1[i]
			}
		}
	}
	return -1
}

/* binary search solution
func getCommon(nums1 []int, nums2 []int) int {
	l, r := 0, len(nums2)-1
	pos := -1
	for i := range nums1 {
		for l <= r && pos == -1 {
			m := (l + r) / 2
			if nums1[i] == nums2[m] {
				pos = nums2[m]
			} else if nums1[i] < nums2[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return pos
}*/

// two pointers approach
func getCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return -1
}
