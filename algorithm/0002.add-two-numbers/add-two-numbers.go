package main

import "fmt"

// package problem0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var merge, cur *ListNode
	var carry, sum int
	carry = 0
	merge = &ListNode{}
	cur = merge
	for l1 != nil || l2 != nil {
		sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		cur.Next = &ListNode{Val: sum - carry*10}
		cur = cur.Next
	}
	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return merge.Next
}

func main() {
	var l, p, m ListNode
	l.Next = &p
	l.Val = 2
	p.Next = &m
	p.Val = 4
	m.Next = nil
	m.Val = 3
	var n1, n2, n3 ListNode
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = nil
	n1.Val = 5
	n2.Val = 6
	n3.Val = 4
	var l1, l2 *ListNode
	l1 = &l
	l2 = &n1
	merge := addTwoNumbers(l1, l2)
	// fmt.Println("hi\n")
	for ; merge != nil; merge = merge.Next {
		fmt.Printf("%d", merge.Val)
	}
}
