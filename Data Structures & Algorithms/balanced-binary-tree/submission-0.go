/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }

    leftH := height(root.Left)
    rightH := height(root.Right)

    if abs(leftH - rightH) > 1 {
        return false
    }

    return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {
    if node == nil {
        return -1
    }
    left := height(node.Left)
    right := height(node.Right)
    return max(left, right) + 1
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}