func reverseKGroup(head *ListNode, k int) *ListNode {

	var dummy ListNode
	dummy.Next = head

	cur := head

	var pre *ListNode
	pre = &dummy

	cnt := 0

	for cur != nil {
		cnt++
		if cnt == k {
			cnt = 0
			pre = reverseOneGroup(pre, cur.Next)
			cur = pre.Next
			continue
		}
		cur = cur.Next
	}
	return dummy.Next
}

/* reverse nodes in (pre, lastNext) and return original first pre.Next */
func reverseOneGroup(pre, lastNext *ListNode) *ListNode {
	cur := pre.Next
	var prev *ListNode
	p := pre.Next
	for cur != lastNext {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	pre.Next = prev
	p.Next = lastNext
	return p
}