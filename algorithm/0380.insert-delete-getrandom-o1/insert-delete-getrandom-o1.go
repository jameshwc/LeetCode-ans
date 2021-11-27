type RandomizedSet struct {
	index map[int]int
	vals  []int
	top   int
}

func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), make([]int, 20000), 0}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.index[val]
	if ok {
		return false
	}
	this.index[val] = this.top
	this.vals[this.top] = val
	this.top++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	id, ok := this.index[val]
	if !ok {
		return false
	}
	if this.vals[this.top-1] != val {
		this.vals[this.top-1], this.vals[id] = this.vals[id], this.vals[this.top-1]
		this.index[this.vals[id]] = id
	}
	delete(this.index, val)
	this.top--
	return true
}

func (this *RandomizedSet) GetRandom() int {
	r := rand.Intn(this.top)
	return this.vals[r]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
