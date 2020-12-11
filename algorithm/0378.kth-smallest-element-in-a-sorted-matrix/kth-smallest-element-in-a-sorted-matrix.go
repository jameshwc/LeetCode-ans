package main

import "fmt"

type Heap struct {
	h         []int
	len       int
	cap       int
	isMinHeap bool
}

func NewHeap(cap int, isMinHeap bool) *Heap {
	var h Heap
	h.h = make([]int, cap)
	h.len = 0
	h.cap = cap
	h.isMinHeap = isMinHeap
	return &h
}
func (h *Heap) Insert(k int) {
	if h.len == h.cap {
		panic("heap is full!")
	}
	h.h[h.len] = k
	h.len++
	h.up(k)
}

func (h *Heap) IsFull() bool {
	return h.len == h.cap
}

func (h *Heap) IsEmpty() bool {
	return h.len == 0
}

func (h *Heap) Top() int {
	if len(h.h) == 0 {
		panic("Heap is empty!")
	}
	return h.h[0]
}
func (h *Heap) Pop() int {
	target := h.h[0]
	h.h[0] = h.h[h.len-1]
	h.len--
	h.down(0)
	return target
}

func (h *Heap) TopK(k int) []int {
	if k <= h.len {
		return h.h[:k]
	}
	return nil
}

func (h *Heap) up(n int) {
	idx := h.len - 1
	for {
		parentIdx := (idx+1)/2 - 1
		if parentIdx < 0 {
			break
		}
		if (h.isMinHeap && h.h[parentIdx] > h.h[idx]) || (!h.isMinHeap && h.h[parentIdx] < h.h[idx]) {
			h.h[parentIdx], h.h[idx] = h.h[idx], h.h[parentIdx]
			idx = parentIdx
		} else {
			break
		}
	}
}

//TODO: isMinHeap / isMaxHeap
func (h *Heap) down(idx int) {
	length := h.len
	for {
		left, right := (idx+1)*2-1, (idx+1)*2
		swapIdx := idx
		if right < length {
			if (h.isMinHeap && h.h[right] < h.h[idx]) || (!h.isMinHeap && h.h[right] > h.h[idx]) {
				if (h.isMinHeap && h.h[right] < h.h[left]) || (!h.isMinHeap && h.h[right] > h.h[left]) {
					swapIdx = right
				} else {
					swapIdx = left
				}
			} else if (h.isMinHeap && h.h[left] < h.h[idx]) || (!h.isMinHeap && h.h[left] > h.h[idx]) {
				swapIdx = left
			}
		} else if left < length {
			if (h.isMinHeap && h.h[left] < h.h[idx]) || (!h.isMinHeap && h.h[left] > h.h[idx]) {
				swapIdx = left
			}
		}
		if swapIdx == idx {
			break
		}
		h.h[idx], h.h[swapIdx] = h.h[swapIdx], h.h[idx]
		idx = swapIdx
	}
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	h := NewHeap(k+1, false)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			h.Insert(matrix[i][j])
			if h.IsFull() {
				h.Pop()
			}
		}
	}
	return h.Top()
}

func upperBound(arr []int, n int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func kthSmallestRevisit(matrix [][]int, k int) int {
	n := len(matrix)
	left := matrix[0][0]
	right := matrix[n-1][n-1]
	for left < right {
		mid := (left + right) / 2
		cnt := 0
		for i := range matrix {
			cnt += upperBound(matrix[i], mid)
		}
		if cnt < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	fmt.Println(kthSmallestRevisit([][]int{
		{1, 2},
		{12, 100},
	}, 3))
}
