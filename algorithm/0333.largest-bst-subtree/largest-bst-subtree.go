/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func largestBSTSubtree(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int, int)
	dfs = func(r *TreeNode) (int, int, int) {
		if r == nil {
			return -(1 << 15), 1 << 15, 0
		}
		leftMin, leftMax, leftNodes := dfs(r.Left)
		rightMin, rightMax, rightNodes := dfs(r.Right)
		if leftMax < r.Val && rightMin > r.Val {
			return leftMin, rightMax, leftNodes + rightNodes + 1
		}
		return -(1 << 15), 1 << 15, max(leftNodes, rightNodes)
	}
	_, _, nums := dfs(root)
	return nums
}