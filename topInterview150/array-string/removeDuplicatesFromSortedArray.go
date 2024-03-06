package main

import "sort"

// pongo le occorrenze delle variabili in una mappa
// pongo le variabili in un'array di supporto
// ordino l'array di supporto, e riassegno gli elementi
// dell'array principali con gli elementi dell'array di supporto
func removeDuplicates(nums []int) int {
	m := make(map[int]bool)
	var supp []int
	for i := range nums {
		if ok := m[nums[i]]; !ok {
			supp = append(supp, nums[i])
		}
		m[nums[i]] = true
	}
	sort.Ints(supp)
	for i := range supp {
		nums[i] = supp[i]
	}
	return len(supp)
}

// abbiamo due indici l e r, tutti e due puntano sul secondo
// elemento dell'array all'inizio. L' indice r avanza a ogni
// iterazione, l avanza quando troviamo i nuovi elementi.
// ad ogni iterazione controlliamo se l'elemento puntato da r e
// il suo predecessore sono diversi, se lo sono poniamo l'elemento
// di indice r in l. (la fa da "segnaposto")
func removeDuplicates(nums []int) int {
	l := 1
	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[r-1] {
			nums[l] = nums[r]
			l++
		}
	}
	return l
}
