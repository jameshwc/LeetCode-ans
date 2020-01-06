package problem0101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkRecursion(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val &&
		checkRecursion(left.Left, right.Right) &&
		checkRecursion(left.Right, right.Left)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkRecursion(root.Left, root.Right)
}
func isSymmetricIteration(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	queue = append(queue, root)
	for len(queue) != 0 {
		var t1, t2 *TreeNode
		t1 = queue[0]
		t2 = queue[1]
		queue = queue[2:]
		if t1 == nil && t2 == nil {
			continue
		} else if t1 == nil || t2 == nil || t1.Val != t2.Val {
			return false
		}
		queue = append(queue, t1.Right)
		queue = append(queue, t2.Left)
		queue = append(queue, t1.Left)
		queue = append(queue, t2.Right)
	}
	return true
}
