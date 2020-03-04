/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/* Sol 1: limit min and max */

func check(root *TreeNode, low, high int) bool {
	if root == nil {
		return true
	}
	if root.Val <= low || root.Val >= high {
		return false
	}
	return check(root.Left, low, root.Val) && check(root.Right, root.Val, high)
}
func isValidBST(root *TreeNode) bool {
	return check(root, -1<<32, 1<<32)
}

/* Sol 2: inorder */

var order []int

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	order = append(order, root.Val)
	inorder(root.Right)
}
func isValidBST(root *TreeNode) bool {
	order = nil
	inorder(root)
	for i := range order {
		if i == 0 {
			continue
		}
		if order[i] <= order[i-1] {
			return false
		}
	}
	return true
}
