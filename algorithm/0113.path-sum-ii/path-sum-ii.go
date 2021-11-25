/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int

	var dfs func(*TreeNode, []int, int)
	dfs = func(r *TreeNode, path []int, sum int) {
		if r == nil {
			return
		}
		if r.Left == nil && r.Right == nil {
			if sum+r.Val == targetSum {
				cp := make([]int, len(path)+1)
				copy(cp, append(path, r.Val))
				ans = append(ans, cp)
			}
			return
		}
		dfs(r.Left, append(path, r.Val), sum+r.Val)
		dfs(r.Right, append(path, r.Val), sum+r.Val)
	}

	dfs(root, []int{}, 0)
	return ans
}