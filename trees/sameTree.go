package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var vis1, vis2 []int
	inorder(p, &vis1)
	inorder(q, &vis2)
	if len(vis1) != len(vis2) {
		return false
	}
	for i := range vis1 {
		if vis1[i] != vis2[i] {
			return false
		}
	}
	return true
}

func inorder(r *TreeNode, visited *[]int) {
	if r == nil {
		return
	}
	if r.Left == nil {
		*visited = append(*visited, 0)
	}
	if r.Right == nil {
		*visited = append(*visited, 0)
	}
	inorder(r.Left, visited)
	*visited = append(*visited, r.Val)
	inorder(r.Right, visited)
}
