/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func invertTree(root *TreeNode) *TreeNode {
    // Base case: If the root is nil, there's nothing to invert.
    if root == nil {
        return nil
    }

    // Recursively invert the left subtree
    // After this call, root.Left will point to the inverted version of its original left subtree
    invertTree(root.Left)

    // Recursively invert the right subtree
    // After this call, root.Right will point to the inverted version of its original right subtree
    invertTree(root.Right)

    // Swap the left and right children of the current node
    // This is the crucial step that inverts the current node's children.
    root.Left, root.Right = root.Right, root.Left

    // Return the root of the now inverted tree
    return root
}