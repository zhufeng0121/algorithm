/**
 leetcode 98 Validate Binary Search Tree

 Given the root of a binary tree, determine if it is a
 valid binary search tree (BST).

 A valid BST is defined as follows:

 The left subtree of a node contains only nodes with keys less than the node's key.
 The right subtree of a node contains only nodes with keys greater than the node's key.
 Both the left and right subtrees must also be binary search trees.

 */
package main

import "math"

func isValidBST(root *TreeNode) bool {
	return helperBST(root,math.MinInt64,math.MaxInt64)

}
func helperBST(root *TreeNode,lower,upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helperBST(root.Left,lower,root.Val) && helperBST(root.Right,root.Val,upper)

}

//通过BST中序遍历有序的性质，检查一棵树是否是BST
func isValidBSTI(root *TreeNode) bool {
	var stack []*TreeNode
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) -1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right

	}
	return true

}
