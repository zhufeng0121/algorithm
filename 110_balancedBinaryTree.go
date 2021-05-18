/**
 leetcode 110 Balanced Binary Tree

 Given a binary tree, determine if it is height-balanced.

 For this problem, a height-balanced binary tree is defined as:

 a binary tree in which the left and right subtrees of every node
 differ in height by no more than 1.

*/
package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if absUtil(maxDepth(root.Left) - maxDepth(root.Right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func absUtil(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
