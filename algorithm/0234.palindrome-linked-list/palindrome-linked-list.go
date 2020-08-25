/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// find middle of the list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// reverse the list from the middle to end
	var prev *ListNode
	for slow != nil {
		temp := slow
		slow = slow.Next
		temp.Next = prev
		prev = temp
	}
	// compare left half and right half
	for prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head = head.Next
		prev = prev.Next
	}
	return true
}