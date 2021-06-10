/**
 leetcode 404 Sum of Left Leaves

Given the root of a binary tree, return the sum of all left leaves.

*/
package main
//left leaves  左叶子节点 不是左节点
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil{
		sum += root.Left.Val
	}
	sum = sum + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	return sum

}

//可以考虑其他的做法
