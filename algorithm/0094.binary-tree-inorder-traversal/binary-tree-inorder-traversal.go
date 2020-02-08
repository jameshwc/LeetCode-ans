/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, ans)
	*ans = append(*ans, root.Val)
	dfs(root.Right, ans)
}
func inorderTraversal(root *TreeNode) []int {
	var ans []int
	dfs(root, &ans)
	return ans
}
func iterInorderTraversal(root *TreeNode) []int {
	var ans []int
	var stack [512]*TreeNode
	top := 0
	cur := root
	for cur != nil || top != 0 {
		if cur {
			stack[top] = cur
			top++
			cur = cur.Left
		} else {
			top--
			cur = stack[top].Right
			ans = append(ans, stack[top].Val)
		}
	}
	return ans
}