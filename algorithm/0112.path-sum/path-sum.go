func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	var dfs func(*TreeNode, int) bool
	dfs = func(r *TreeNode, remain int) bool {
		if r == nil {
			return false
		}
		remain -= r.Val
		if r.Left == nil && r.Right == nil {
			return remain == 0
		} else if r.Left == nil {
			return
		}
		return dfs(r.Left, remain) || dfs(r.Right, remain)
	}
	return dfs(root, targetSum)
}