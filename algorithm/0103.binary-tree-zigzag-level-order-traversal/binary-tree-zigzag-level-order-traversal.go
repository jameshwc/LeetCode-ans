func zigzagLevelOrder(root *TreeNode) [][]int {

	var nodes [][]int
	var ans [][]int
	var curLayer []*TreeNode
	var nextLayer []*TreeNode

	if root == nil {
		return ans
	}

	curLayer = append(curLayer, root)
	nodes = append(nodes, append([]int{}, root.Val))

	height := 0

	for len(curLayer) > 0 {

		node := curLayer[0]

		if node.Left != nil {
			nextLayer = append(nextLayer, node.Left)
		}
		if node.Right != nil {
			nextLayer = append(nextLayer, node.Right)
		}

		curLayer = curLayer[1:]
		if len(curLayer) == 0 {
			var nodeVals []int
			for i := range nextLayer {
				nodeVals = append(nodeVals, nextLayer[i].Val)
			}
			nodes = append(nodes, nodeVals)
			curLayer = nextLayer
			nextLayer = make([]*TreeNode, 0)
			height++
		}
	}

	for i := range nodes {
		if i%2 == 0 {
			ans = append(ans, nodes[i])
		} else {
			var rev []int
			for j := len(nodes[i]) - 1; j > -1; j-- {
				rev = append(rev, nodes[i][j])
			}
			ans = append(ans, rev)
		}
	}
	return ans[:len(ans)-1]
}