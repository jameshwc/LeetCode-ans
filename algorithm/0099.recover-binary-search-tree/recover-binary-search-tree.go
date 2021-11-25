package main

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func recoverTree(root *TreeNode) {
	var traverse func(*TreeNode)
	var nodes []*TreeNode
	var vals []int
	traverse = func(r *TreeNode) {
		if r == nil {
			return
		}
		traverse(r.Left)
		nodes = append(nodes, r)
		vals = append(vals, r.Val)
		traverse(r.Right)
	}
	traverse(root)
	sort.Ints(vals)
	for i := 0; i < len(nodes); i++ {
		nodes[i].Val = vals[i]
	}
}
