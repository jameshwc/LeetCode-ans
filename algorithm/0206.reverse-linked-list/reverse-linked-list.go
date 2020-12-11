// 1->2->3->4->5->nil
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	next := head.Next
	head.Next = nil
	for fast != nil && next != nil {
		origin := slow
		fast = next.Next
		slow = next
		slow.Next = origin
		if fast != nil {
			next = fast.Next
			fast.Next = slow
			slow = fast
		}
	}
	return slow
}

func reverseListSimple(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}