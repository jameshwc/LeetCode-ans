/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head, head
	for {
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		if fast == nil {
			return false
		}
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
}
