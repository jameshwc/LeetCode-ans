package problem0101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkRecursion(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val &&
		checkRecursion(left.Left, right.Right) &&
		checkRecursion(left.Right, right.Left)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkRecursion(root.Left, root.Right)
}
