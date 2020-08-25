/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow, pre := head, head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		fast = fast.Next.Next
		slow = slow.Next
	}
	pre.Next = nil
	return merge(sortList(head), sortList(slow))
}

func merge(left, right *ListNode) *ListNode {
	var head *ListNode
	cur := &ListNode{}
	head = cur
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}
	return head.Next
}