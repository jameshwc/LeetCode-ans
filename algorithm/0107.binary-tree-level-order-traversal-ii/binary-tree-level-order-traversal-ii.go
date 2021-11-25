/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	nodesWithHeight := make([][]int, 128)
	var traverse func(*TreeNode, int)
	traverse = func(r *TreeNode, h int) {
		if r == nil {
			return
		}
		traverse(r.Left, h+1)
		nodesWithHeight[h] = append(nodesWithHeight[h], r.Val)
		traverse(r.Right, h+1)
	}
	traverse(root, 0)
	var ans [][]int
	for i := len(nodesWithHeight) - 1; i > -1; i-- {
		if len(nodesWithHeight[i]) == 0 {
			continue
		}
		ans = append(ans, nodesWithHeight[i])
	}
}