/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var inorder func(*TreeNode)
	cnt, isFound, ans := 0, false, 0
	inorder = func(r *TreeNode) {
		if r == nil || isFound {
			return
		}
		inorder(r.Left)
		cnt++
		if cnt == k {
			isFound = true
			ans = r.Val
		}
		inorder(r.Right)
	}
	inorder(root)
	return ans
}