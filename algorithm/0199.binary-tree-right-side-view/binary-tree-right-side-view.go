/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var ans []int
	var queue []*TreeNode
	var values [][]int
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		vals := make([]int, n)
		for i := 0; i < n; i++ {
			vals[i] = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		values = append(values, vals)
		queue = queue[n:]
	}
	for i := range values {
		ans = append(ans, values[i][len(values[i])-1])
	}
	return ans
}