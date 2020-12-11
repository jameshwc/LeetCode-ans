/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var isLeftFind, isRightFind bool
	var leftPath, rightPath []*TreeNode
	var findPath func(*TreeNode, []*TreeNode)
	findPath = func(r *TreeNode, path []*TreeNode) {
		if r == nil || (isLeftFind && isRightFind) {
			return
		}
		path = append(path, r)
		if r == p {
			isLeftFind = true
			leftPath = path
		} else if r == q {
			isRightFind = true
			rightPath = path
		}
		n := len(path)
		findPath(r.Left, path)
		path = path[:n]
		findPath(r.Right, path)
		path = path[:n]
	}
	var path []*TreeNode
	findPath(root, path)
	min := len(leftPath)
	if min > len(rightPath) {
		min = len(rightPath)
	}
	i := 1
	for ; i < min; i++ {
		if leftPath[i] != rightPath[i] {
			return leftPath[i-1]
		}
	}
	return leftPath[i-1]
}

func sol2(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}
	left := sol2(root.Left, p, q)
	right := sol2(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}