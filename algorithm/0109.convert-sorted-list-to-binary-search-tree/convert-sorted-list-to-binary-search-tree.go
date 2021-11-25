/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	cur := head
	n := 0
	for cur != nil {
		n++
		cur = cur.Next
	}

	var inorder func(int, int) *TreeNode
	inorder = func(left, right int) *TreeNode {

		if left > right {
			return nil
		}

		var root TreeNode

		mid := (left + right) >> 1

		root.Left = inorder(left, mid)

		root.Val = head.Val
		head = head.Next

		root.Right = inorder(mid+1, right)

		return root
	}

	return inorder(0, n)
}