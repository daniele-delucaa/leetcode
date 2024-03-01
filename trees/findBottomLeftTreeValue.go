package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// fai una BFS sull'albero (implementata grazie l'utilizzo di una coda), prima visito il figlio dx, poi il figlio sx
// l'ultimo nodo visitato dalla BFS sarà il nodo da restituire

type queue struct {
	items []TreeNode
}

func (q *queue) Enqueue(n TreeNode) {
	q.items = append(q.items, n)
}

func (q *queue) Dequeue() (TreeNode, bool) {
	if q.isEmpty() {
		return TreeNode{}, false
	}
	v := q.items[0]
	q.items = q.items[1:]
	return v, true
}

func (q *queue) isEmpty() bool {
	return len(q.items) == 0
}

func findBottomLeftValue(root *TreeNode) int {
	var curr TreeNode
	q := queue{}

	n := root
	q.Enqueue(*n)
	for !q.isEmpty() {
		curr, _ = q.Dequeue()
		fmt.Println(curr)
		if curr.Right != nil {
			q.Enqueue(*curr.Right)
		}
		if curr.Left != nil {
			q.Enqueue(*curr.Left)
		}
	}
	return curr.Val
}

func main() {
	t := TreeNode{2, nil, nil}
	t.Left = &TreeNode{1, nil, nil}
	t.Right = &TreeNode{3, nil, nil}
	fmt.Println(findBottomLeftValue(&t))
}
