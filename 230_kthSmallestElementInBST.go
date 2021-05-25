/**
 leetcode 230 Kth Smallest Element in a BST

 Given the root of a binary search tree, and an integer k, return the
 kth (1-indexed) smallest element in the tree.


 */
package main

//先
func kthSmallestI(root *TreeNode, k int) int {
	return inOrder(root)[k-1]

}

func inOrder(root * TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int,0)
	result = append(result,inOrder(root.Left)...)
	result = append(result, root.Val)
	result = append(result,inOrder(root.Right)...)
	return result
}


//计算子树节点的数量
func kthSmallestII(root *TreeNode, k int) int {
	count := countNodes(root.Left)
	if k <= count {
		return kthSmallestII(root.Left,k)
	} else if k > count + 1 {
		return kthSmallestII(root.Right,k-count-1)
	}
	return root.Val

}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
