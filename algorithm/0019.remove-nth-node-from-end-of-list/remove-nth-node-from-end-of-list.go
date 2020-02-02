package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	t := head
	for cur != nil {
		if n < 0 {
			t = t.Next
		}
		n--
		cur = cur.Next
	}
	if n < 0 {
		t.Next = t.Next.Next
		return head
	}
	return head.Next
}
func main() {

}
