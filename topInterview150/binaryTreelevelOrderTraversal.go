package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	currLevel := []*TreeNode{root}

	for len(currLevel) > 0 {
		var nextLevel []*TreeNode
		res = append(res, extractNodeValues(currLevel))
		for _, n := range currLevel {
			if n.Left != nil {
				nextLevel = append(nextLevel, n.Left)
			}
			if n.Right != nil {
				nextLevel = append(nextLevel, n.Right)
			}
		}
		currLevel = nextLevel
	}
	return res
}

func extractNodeValues(nodes []*TreeNode) []int {
	var res []int
	for _, n := range nodes {
		res = append(res, n.Val)
	}
	return res
}
