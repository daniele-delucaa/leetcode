package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// soluzione ricorsiva
// un'altro approccio poteva essere usare una BFS e una variabile booleana
// ogni volta che si aggiungono dei nodi alla coda settiamo a true
// la variabile, quando è true incrementiamo il contatore e
// risettiamo a false la var booleana
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
