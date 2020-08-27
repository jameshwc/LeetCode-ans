package heap

type Heap struct {
	h   []int
	len int
	cap int
}

func NewHeap(cap int) *Heap {
	var h Heap
	h.h = make([]int, cap)
	h.len = 0
	h.cap = cap
	return &h
}
func (h *Heap) Insert(n, k int) {
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
		if h.h[parentIdx] > h.h[idx] {
			h.h[parentIdx], h.h[idx] = h.h[idx], h.h[parentIdx]
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
		idx = swapIdx
	}
}
