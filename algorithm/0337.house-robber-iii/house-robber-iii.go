/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int)
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		leftRob, leftUnrob := dfs(root.Left)
		rightRob, rightUnrob := dfs(root.Right)
		maxLeft, maxRight := max(leftRob, leftUnrob), max(rightRob, rightUnrob)
		return leftUnrob + rightUnrob + root.Val, maxLeft + maxRight
	}
	return max(dfs(root))
}