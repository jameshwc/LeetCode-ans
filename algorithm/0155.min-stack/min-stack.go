package main

type MinStack struct {
	data     [8192]int
	minStack [8192]int
	top      int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var ms MinStack
	ms.top = 0
	return ms
}

func (this *MinStack) Push(x int) {
	this.data[this.top] = x
	if this.top == 0 || this.minStack[this.top-1] > x {
		this.minStack[this.top] = x
	} else {
		this.minStack[this.top] = this.minStack[this.top-1]
	}
	this.top++
}

func (this *MinStack) Pop() {
	this.top--
}

func (this *MinStack) Top() int {
	return this.data[this.top-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[this.top-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(3)
	obj.Pop()
}
