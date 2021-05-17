/**
 leetcode 107 Binary Tree Level Order Traversal II

 Given the root of a binary tree, return the bottom-up level order traversal of
 its nodes' values. (i.e., from left to right, level by level from leaf to root).
 */
package main

func levelOrderBottom(root *TreeNode) [][]int {
	queue := make([]*TreeNode,0)
	var result [][]int
	var reverse [][]int
	if root == nil {
		return nil
	}
	//入队 root节点 第一层
	queue = append(queue,root)
	for len(queue) != 0 {
		level := make([]int,0)
		size := len(queue)
		for i := 0;i < size;i++ {
			temp := queue[0]
			queue = queue[1:]
			if temp.Left != nil {
				queue = append(queue,temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue,temp.Right)
			}
			level = append(level, temp.Val)

		}
		result = append(result,level)

	}
	if len(result) == 0 || len(result) == 1 {
		return result
	}
	for i := len(result) - 1;i >= 0;i-- {
		reverse = append(reverse,result[i])
	}

	return reverse

}
