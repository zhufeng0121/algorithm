/**
 leetcode 144 Binary Tree Preorder Traversal

 Given the root of a binary tree, return the preorder traversal of
 its nodes' values.


 */
package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
//递归做法
func preorderTraversal(root *TreeNode) []int {
	preorder := make([]int,0)
	if root == nil {
		return []int{}
	}
	preorder = append(preorder,root.Val)
	preorder = append(preorder,preorderTraversal(root.Left)...)
	preorder = append(preorder,preorderTraversal(root.Right)...)
	return preorder

}
//熟悉一下别的做法