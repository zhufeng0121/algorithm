/**
 leetcode 101 Symmetric Tree

 Given the root of a binary tree, check whether it is a mirror of
 itself (i.e., symmetric around its center).
 */
package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left,root.Right)

}

func symmetric(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	//注意这个写法
	return symmetric(t1.Left,t2.Right)  && symmetric(t1.Right,t2.Left)
}




