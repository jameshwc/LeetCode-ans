package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var root ListNode
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	prev := &root
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = cur.Next.Next
		next.Next = cur
		prev.Next = next
		prev = cur
		cur = cur.Next
	}
	return root.Next
}

func main() {

}
