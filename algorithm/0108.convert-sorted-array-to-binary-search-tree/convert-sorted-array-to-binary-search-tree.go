/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	idx := 0

	var inorder func(int, int) *TreeNode
	inorder = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}

		var root TreeNode

		mid := (l + r) >> 1
		root.Left = inorder(l, mid-1)

		root.Val = nums[idx]
		idx++

		root.Right = inorder(mid+1, r)
		return &root
	}

	return inorder(0, len(nums)-1)
}