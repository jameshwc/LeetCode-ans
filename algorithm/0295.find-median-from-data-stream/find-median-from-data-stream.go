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

func (h *Heap) Len() int {
	return h.len
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

type MedianFinder struct {
	leftHeap  *Heap
	rightHeap *Heap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	var m MedianFinder
	m.leftHeap = NewHeap(1<<20, false)
	m.rightHeap = NewHeap(1<<20, true)
	return m
}

func (this *MedianFinder) AddNum(num int) {
	this.leftHeap.Insert(num)
	this.rightHeap.Insert(this.leftHeap.Pop())
	if this.leftHeap.Len() < this.rightHeap.Len() {
		this.leftHeap.Insert(this.rightHeap.Pop())
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.leftHeap.Len() > this.rightHeap.Len() {
		return float64(this.leftHeap.Top())
	}
	return float64(this.leftHeap.Top()+this.rightHeap.Top()) * 0.5
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

// 7 3 19 2 8 2 4 5 6
// 7
// 3 7
// [3] [7,19]
// []
