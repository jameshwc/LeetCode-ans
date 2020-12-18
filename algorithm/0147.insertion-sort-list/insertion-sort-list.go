/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	cur := head
	var dummy ListNode
	for cur != nil {
		prev := &dummy
		for prev.Next != nil && prev.Next.Val < cur.Val {
			prev = prev.Next
		}
		next := cur.Next
		cur.Next = prev.Next
		prev.Next = cur
		cur = next
	}
	return dummy.Next
}