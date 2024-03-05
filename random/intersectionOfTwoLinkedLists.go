package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	occA := make(map[*ListNode]bool)
	occB := make(map[*ListNode]bool)
	currA := headA
	currB := headB
	// i due for iniziali servono nel caso le due liste sono uguali
	// (hanno gli stessi nodi e lunghezza) e ad un certo punto si
	// intersecano, infstti i due if iniziali del ciclo principale
	// del programma non rileverebbe le occorrenze, perche la mappa
	// verrebbe alimentata solo dopo l'iterazione
	for currA != nil {
		occA[currA] = true
		currA = currA.Next
	}
	for currB != nil {
		occB[currB] = true
		currB = currB.Next
	}
	currA = headA
	currB = headB
	for currA != nil || currB != nil {
		if currB != nil && occA[currB] == true {
			return currB
		}
		if currA != nil && occB[currA] == true {
			return currA
		}
		if currA != nil {
			currA = currA.Next
		}
		if currB != nil {
			currB = currB.Next
		}
	}
	return nil
}
