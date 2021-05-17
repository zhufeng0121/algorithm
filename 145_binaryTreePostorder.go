/**
 leetcode 145 Binary Tree Postorder Traversal

 Given the root of a binary tree, return the postorder traversal of its nodes' values.
 */
package main

func postorderTraversal(root *TreeNode) []int {
	postorder := make([]int,0)
	if root == nil {
		return []int{}
	}
	postorder = append(postorder,postorderTraversal(root.Left)...)
	postorder = append(postorder,postorderTraversal(root.Right)...)
	postorder = append(postorder,root.Val)
	return postorder

}