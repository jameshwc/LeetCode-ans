package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// type TreeNode struct {
// Val   int
// Left  *TreeNode
// Right *TreeNode
// }

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res []string
	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			res = append(res, "null")
			return
		}
		res = append(res, strconv.Itoa(r.Val))
		dfs(r.Left)
		dfs(r.Right)
	}
	dfs(root)
	return strings.Join(res, " ")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodesVal := strings.Split(data, " ")
	top := 0
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if nodesVal[top] == "null" {
			top++
			return nil
		}
		var r TreeNode
		r.Val, _ = strconv.Atoi(nodesVal[top])
		top++
		r.Left = dfs()
		r.Right = dfs()
		return &r
	}
	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
