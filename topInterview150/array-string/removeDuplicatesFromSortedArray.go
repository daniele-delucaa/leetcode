package main

import "sort"

// pongo le occorrenze delle variabili in una mappa
// pongo le variabili in un'array di supporto
// ordino l'array di supporto, e gli elementi dell'array di
// supporto li metto in quello principale
func removeDuplicates(nums []int) int {
	m := make(map[int]int)
	var supp []int
	for i := range nums {
		m[nums[i]]++
	}
	for k := range m {
		supp = append(supp, k)
	}
	sort.Ints(supp)
	for i := range supp {
		nums[i] = supp[i]
	}
	return len(supp)
}
