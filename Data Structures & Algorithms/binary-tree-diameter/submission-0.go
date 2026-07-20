/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0

	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftDepth := depth(node.Left)
		rightDepth := depth(node.Right)

		diameter = max(diameter, leftDepth+rightDepth)

		return max(leftDepth, rightDepth) + 1
	}

	depth(root)

	return diameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
