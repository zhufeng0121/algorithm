/**
 leetcode 226 Invert Binary Tree

 Given the root of a binary tree, invert the tree, and return its root.

*/
package main

//翻转二叉树 ，那这样的话，判断二叉树是否互为镜像，也是可以的了
//居然代码完全一样
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root

}
