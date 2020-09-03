/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var dfs func(*TreeNode, int) int
	dfs = func(root *TreeNode, cur int) int {
		if root == nil {
			return 0
		}
		cur += root.Val
		cnt := 0
		if cur == sum {
			cnt++
		}
		return cnt + dfs(root.Left, cur) + dfs(root.Right, cur)
	}
	return dfs(root, 0) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}