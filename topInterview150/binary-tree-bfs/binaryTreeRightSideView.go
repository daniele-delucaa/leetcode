package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var allNodes [][]*TreeNode
	currLevel := []*TreeNode{root}
	for len(currLevel) > 0 {
		var nextLevel []*TreeNode
		allNodes = append(allNodes, currLevel)
		for _, n := range currLevel {
			if n.Right != nil {
				nextLevel = append(nextLevel, n.Right)
			}
			if n.Left != nil {
				nextLevel = append(nextLevel, n.Left)
			}
		}
		currLevel = nextLevel
	}
	for _, row := range allNodes {
		for i, n := range row {
			if i == 0 {
				res = append(res, n.Val)
			}
		}
	}
	return res
}
