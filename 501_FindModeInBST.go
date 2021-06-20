/**
 leetcode 501. Find Mode In Binary Search Tree

 Given the root of a binary search tree (BST) with duplicates,
 return all the mode(s) (i.e., the most frequently occurred element) in it.

 If the tree has more than one mode, return them in any order.

 Assume a BST is defined as follows:

 The left subtree of a node contains only nodes with keys less than or equal to the node's key.
 The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
 Both the left and right subtrees must also be binary search trees.


 */
package main

import "math"

//DFS 中序遍历 + map统计 (通过了，但是较慢)
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nodes := make(map[int]int)
	stack := make([]*TreeNode,0)
	res := make([]int,0)
	for len(stack) != 0 || root != nil{
		if root != nil {
			stack = append(stack,root)
			root = root.Left
		} else {
			root = stack[len(stack) -1]
			stack = stack[:len(stack) -1]
			nodes[root.Val]++
			root = root.Right
		}
	}
	max := math.MinInt32
	for _, v := range nodes {
		if v >= max {
			max  = v
		}
	}
	for i, v := range nodes {
		if v == max {
			res = append(res,i)
		}
	}
	return res

}
//是否可以考虑采用递归 ，待优化

func findModeRec(root *TreeNode) []int {
	if root == nil {
		return nil
	}
}

