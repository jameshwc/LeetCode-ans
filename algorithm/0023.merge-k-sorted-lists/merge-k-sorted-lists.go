package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Heaper struct {
	Items    []interface{}
	LessFunc func(int, int) bool
}

// Len is the number of elements in the collection.
func (h Heaper) Len() int {
	return len(h.Items)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (h Heaper) Less(i int, j int) bool {
	return h.LessFunc(i, j)
}

// Swap swaps the elements with indexes i and j.
func (h *Heaper) Swap(i int, j int) {
	h.Items[i], h.Items[j] = h.Items[j], h.Items[i]
}

func (h *Heaper) Push(x interface{}) {
	h.Items = append(h.Items, x)
}

func (h *Heaper) Pop() interface{} {
	item := h.Items[len(h.Items)-1]
	h.Items = h.Items[:len(h.Items)-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	var pq Heaper
	pq.LessFunc = func(i, j int) bool {
		return pq.Items[i].(*ListNode).Val < pq.Items[j].(*ListNode).Val
	}
	var head ListNode
	cur := &head
	for i := range lists {
		if lists[i] != nil {
			heap.Push(&pq, lists[i])
		}
	}
	for len(pq.Items) != 0 {
		node := heap.Pop(&pq).(*ListNode)
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
		cur.Next = node
		cur = cur.Next
	}
	return head.Next
}
