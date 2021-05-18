/**
 leetcode 108 Convert Sorted Array to Binary Search Tree

 Given an integer array nums where the elements are sorted in ascending
 order, convert it to a height-balanced binary search tree.

 A height-balanced binary tree is a binary tree in which the depth of the
 two subtrees of every node never differs by more than one.
 */
package main

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[len(nums) / 2]}
	root.Left = sortedArrayToBST(nums[:len(nums)/2])
	root.Right = sortedArrayToBST(nums[len(nums)/2 + 1:])
	return root

}
