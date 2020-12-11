func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left // preorder or postoder is fine
	return root
}

func invertTreeIterate(root *TreeNode) *TreeNode {
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == nil {
			continue
		}
		cur.Left, cur.Right = cur.Right, cur.Left
		queue = append(queue, cur.Left)
		queue = append(queue, cur.Right)
	}
	return root
}