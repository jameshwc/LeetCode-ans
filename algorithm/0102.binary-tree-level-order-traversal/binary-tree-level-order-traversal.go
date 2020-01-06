package problem0102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans = make([][]int, 0)
	var t1 = make([]*TreeNode, 0)
	var t2 = make([]*TreeNode, 0)
	var curLayer = make([]int, 0)
	t1 = append(t1, root)
	for len(t1) != 0 {
		curLayer = nil
		for i, l := 0, len(t1); i < l; i++ {
			curLayer = append(curLayer, t1[i].Val)
			if t1[i].Left != nil {
				t2 = append(t2, t1[i].Left)
			}
			if t1[i].Right != nil {
				t2 = append(t2, t1[i].Right)
			}
		}
		t1 = t2
		t2 = nil
		ans = append(ans, curLayer)
	}
	return ans
}
