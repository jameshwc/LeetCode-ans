package main

type Heap struct {
	h   []int
	n   []int
	len int
	cap int
}

func NewHeap(cap int) *Heap {
	var h Heap
	h.h = make([]int, cap)
	h.n = make([]int, cap)
	h.len = 0
	h.cap = cap
	return &h
}
func (h *Heap) Insert(n, k int) {
	if h.len == h.cap {
		panic("heap is full!")
	}
	h.h[h.len] = k
	h.n[h.len] = n
	h.len++
	h.up(k)
}

func (h *Heap) IsFull() bool {
	return h.len == h.cap
}

func (h *Heap) Pop() int {
	target := h.h[0]
	h.h[0] = h.h[h.len-1]
	h.n[0] = h.n[h.len-1]
	h.len--
	h.down(0)
	return target
}

func (h *Heap) TopK(k int) []int {
	return h.n[:k]
}

func (h *Heap) up(n int) {
	idx := h.len - 1
	for {
		parentIdx := (idx+1)/2 - 1
		if parentIdx < 0 {
			break
		}
		if h.h[parentIdx] > h.h[idx] {
			h.h[parentIdx], h.h[idx] = h.h[idx], h.h[parentIdx]
			h.n[parentIdx], h.n[idx] = h.n[idx], h.n[parentIdx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *Heap) down(idx int) {
	length := h.len
	for {
		left, right := (idx+1)*2-1, (idx+1)*2
		swapIdx := idx
		if right < length {
			if h.h[right] < h.h[idx] {
				if h.h[right] < h.h[left] {
					swapIdx = right
				} else {
					swapIdx = left
				}
			} else if h.h[left] < h.h[idx] {
				swapIdx = left
			}
		} else if left < length {
			if h.h[left] < h.h[idx] {
				swapIdx = left
			}
		}
		if swapIdx == idx {
			break
		}
		h.h[idx], h.h[swapIdx] = h.h[swapIdx], h.h[idx]
		h.n[idx], h.n[swapIdx] = h.n[swapIdx], h.n[idx]
		idx = swapIdx
	}
}

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	for i := range nums {
		counter[nums[i]]++
	}
	h := NewHeap(k + 1)
	for num, frequency := range counter {
		h.Insert(num, frequency)
		if h.IsFull() {
			h.Pop()
		}
	}
	return h.TopK(k)
}

func quickSelectiveSolution(nums []int, k int) []int {
	counter := make(map[int]int)
	uniqueNums := make([]int, len(nums))
	for i := range nums {
		counter[nums[i]]++
	}
	idx := 0
	for k := range counter {
		uniqueNums[idx] = k
		idx++
	}
	left, right := 0, idx-1
	for left <= right {
		pivot := right
		storeIdx := left
		for i := left; i < right; i++ {
			if counter[uniqueNums[i]] > counter[uniqueNums[pivot]] {
				uniqueNums[i], uniqueNums[storeIdx] = uniqueNums[storeIdx], uniqueNums[i]
				storeIdx++
			}
		}
		uniqueNums[pivot], uniqueNums[storeIdx] = uniqueNums[storeIdx], uniqueNums[pivot]
		if storeIdx == k {
			return uniqueNums[:k]
		} else if storeIdx < k {
			left = storeIdx + 1
		} else {
			right = storeIdx - 1
		}
	}
	return uniqueNums[:k]
}
