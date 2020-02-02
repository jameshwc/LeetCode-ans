package main

import "fmt"

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var merge, cur *ListNode
	merge = &ListNode{}
	cur = merge
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val > l2.Val {
				cur.Next = l2
				l2 = l2.Next
			} else {
				cur.Next = l1
				l1 = l1.Next
			}
			cur = cur.Next
		} else if l1 != nil {
			cur.Next = l1
			break
		} else {
			cur.Next = l2
			break
		}
	}
	return merge.Next
}

func main() {
	var l1, l2, l1_2, l1_3, l2_2, l2_3 ListNode
	l1.Val = 1
	l1.Next = &l1_2
	l1_2.Val = 2
	l1_2.Next = &l1_3
	l1_3.Val = 4
	l1_3.Next = nil
	l2.Val = 1
	l2.Next = &l2_2
	l2_2.Val = 3
	l2_2.Next = &l2_3
	l2_3.Val = 4
	l2_3.Next = nil
	c := mergeTwoLists(&l1, &l2)
	for c != nil {
		fmt.Println(c.Val)
		c = c.Next
	}
}
