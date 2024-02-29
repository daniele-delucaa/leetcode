package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	res := inorder(root)
	return res
}

func inorder(r *TreeNode) int {
	var res int
	var height, maxHeight int
	height++
	if r == nil {
		return res
	}
	if r.Left == nil && height > maxHeight {
		res = r.Val
		maxHeight = height
		height = 0
	}
	inorder(r.Left)
	inorder(r.Right)
	return res
}
