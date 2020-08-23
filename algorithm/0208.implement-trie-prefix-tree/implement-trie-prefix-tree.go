type Trie struct {
	isEnd bool
	nodes [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	current := this
	for _, ch := range word {
		index := ch - 'a'
		if current.nodes[index] == nil {
			current.nodes[index] = &Trie{}
		}
		current = current.nodes[index]
	}
	current.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	current := this
	for _, ch := range word {
		index := ch - 'a'
		if current.nodes[index] == nil {
			return false
		}
		current = current.nodes[index]
	}
	return current.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	current := this
	for _, ch := range prefix {
		index := ch - 'a'
		if current.nodes[index] == nil {
			return false
		}
		current = current.nodes[index]
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */