/**
 leetcode 94 Binary Tree Inorder Traversal

 Given the root of a binary tree, return the inorder traversal of its
 nodes' values.
 */
package main

//递归做法
func inorderTraversal(root *TreeNode) []int {
	inorder := make([]int,0)
	if root == nil {
		return []int{}
	}
	inorder = append(inorder,inorderTraversal(root.Left)...)
	inorder = append(inorder,root.Val)
	inorder = append(inorder,inorderTraversal(root.Right)...)
	return inorder


}
//探索一下是否还有别的做法