/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	Val          int
	NextIterator *BSTIterator
}

func Constructor(root *TreeNode) BSTIterator {
	var inorder func(*TreeNode)
	var dummy BSTIterator
	cur := &dummy
	inorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		inorder(r.Left)
		next := new(BSTIterator)
		next.Val = r.Val
		cur.NextIterator = next
		cur = next
		inorder(r.Right)
	}
	inorder(root)
	return dummy
}

func (this *BSTIterator) Next() int {
	r = this.NextIterator.Val
	this.NextIterator = this.NextIterator.NextIterator
	return r
}

func (this *BSTIterator) HasNext() bool {
	return this.NextIterator == nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

/* * * * * * * * * * * * * * * * * * * * * * * * */

type BSTIterator struct {
	values []int
	idx    int
}

func Constructor(root *TreeNode) BSTIterator {
	var inorder func(*TreeNode)
	var dummy BSTIterator
	inorder = func(r *TreeNode) {
		if r == nil {
			return
		}
		inorder(r.Left)
		dummy.values = append(dummy.values, r.Val)
		inorder(r.Right)
	}
	inorder(root)
	return dummy
}

func (this *BSTIterator) Next() int {
	r := this.values[this.idx]
	this.idx++
	return r
}

func (this *BSTIterator) HasNext() bool {
	return this.idx == len(this.values)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */