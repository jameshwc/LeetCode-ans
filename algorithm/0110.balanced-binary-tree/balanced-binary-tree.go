func abs(p int) int {
	if p > 0 {
		return p
	}
	return p * -1
}

func max(p, q int) int {
	if p > q {
		return p
	}
	return q
}

func dfs(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	left, lbalanced := dfs(root.Left)
	right, rbalanced := dfs(root.Right)

	if lbalanced == false || rbalanced == false || abs(left-right) > 1 {
		return 0, false
	}

	return 1 + max(left, right), true
}

func isBalanced(root *TreeNode) bool {
	_, b := dfs(root)
	return b
}