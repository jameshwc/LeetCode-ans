/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	var dfs func(*TreeNode) int
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	ans := 0
	dfs = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		left := max(dfs(r.Left), 0)
		right := max(dfs(r.Right), 0)
		ans = max(left+right+r.Val, ans)
		return max(left, right) + r.Val
	}
}