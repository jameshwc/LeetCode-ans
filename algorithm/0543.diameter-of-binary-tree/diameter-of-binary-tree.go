/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 1
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		L, R := dfs(root.Left), dfs(root.Right)
		total := L + R + 1
		if total > ans {
			ans = total
		}
		if L > R {
			return L + 1
		}
		return R + 1
	}
	dfs(root)
	return ans - 1
}