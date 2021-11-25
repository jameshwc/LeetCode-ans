/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {

	postTail := len(postorder) - 1
	valuesToPostIndex := make(map[int]int)

	var dfs func(int, int) *TreeNode
	dfs = func(inLeft, inRight int) *TreeNode {
		if inRight < inLeft {
			return nil
		}
		var root TreeNode
		rootIdx := valuesToPostIndex[postorder[postTail]]
		root.Val = inorder[rootIdx]

		postTail--
		root.Right = dfs(rootIdx+1, inRight)
		root.Left = dfs(inLeft, rootIdx-1)
		return &root
	}

	for i := 0; i < len(inorder); i++ {
		valuesToPostIndex[inorder[i]] = i
	}
	return dfs(0, len(inorder)-1)
}