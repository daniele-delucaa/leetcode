package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	var i int = m + n - 1
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
		i--
	}
	for n > 0 {
		nums1[i] = nums2[n-1]
		n--
		i--
	}
}

func main() {
	a1 := []int{2, 0}
	a2 := []int{1}
	merge(a1, 1, a2, 1)
	fmt.Println(a1)
}
