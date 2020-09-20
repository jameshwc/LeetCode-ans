/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	var dfs func(*TreeNode) *TreeNode
	dfs = func(r *TreeNode) *TreeNode {
		if r == nil {
			return
		}
		temp := r.Right
		r.Right = dfs(r.Left)
	}
}