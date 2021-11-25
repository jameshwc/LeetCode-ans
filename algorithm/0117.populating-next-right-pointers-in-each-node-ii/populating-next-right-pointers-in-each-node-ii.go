/**
* Definition for a Node.
* type Node struct {
*     Val int
*     Left *Node
*     Right *Node
*     Next *Node
* }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var queue []*Node
	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {

			node := queue[i]
			if i+1 == n {
				node.Next = nil
			} else {
				node.Next = queue[i+1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
	}
	return root
}