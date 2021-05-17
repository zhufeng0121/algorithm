/**
 leetcode 199 Binary Tree Right Side View

 Given the root of a binary tree, imagine yourself standing on the right side of it,
 return the values of the nodes you can see ordered from top to bottom.

 */
package main

//返回一颗二叉树的右视图，实际上就是返回每层遍历的最右节点而已
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode,0)
	result := make([]int,0)
	queue = append(queue,root)
	for len(queue) != 0 {
		right := 0
		size := len(queue)
		for i := 0;i < size;i++ {
			temp := queue[0]
			queue := queue[1:]
			if temp.Left != nil {
				queue = append(queue,temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue,temp.Right)
			}
			right = temp.Val

		}
		result = append(result,right)
	}
	return result

}

// 有点好奇的是这个答案居然击败了100%
func rightSideViewII(root *TreeNode) []int {
	queue := make([]*TreeNode,0)
	var result []int
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
		result = append(result,level[len(level)-1])

	}
	return result
}

// 有点好奇的是这个答案居然击败了100%
func rightSideViewIII(root *TreeNode) []int {
	queue := make([]*TreeNode,0)
	var result []int
	if root == nil {
		return nil
	}
	//入队 root节点 第一层
	queue = append(queue,root)
	for len(queue) != 0 {
		right := 0
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
			right = temp.Val

		}
		result = append(result,right)

	}
	return result
}
