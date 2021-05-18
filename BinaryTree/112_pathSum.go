/**
 leetcode 112 Path Sum

 Given the root of a binary tree and an integer targetSum,
 return true if the tree has a root-to-leaf path such that
 adding up all the values along the path equals targetSum.

 A leaf is a node with no children.
 */
package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root .Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}
	return hasPathSum(root.Left,targetSum - root.Val) || hasPathSum(root.Right,targetSum - root.Val)
}
