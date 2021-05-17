/**
 leetcode 103 Binary Tree Zigzag Level Order Traversal

 Given the root of a binary tree, return the zigzag level order traversal of its nodes' values.
 (i.e., from left to right, then right to left for the next level and alternate between).
 */
package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode,0)
	var result [][]int
	//将根节点添加置队列之中
	queue = append(queue, root)
	reverse := false
	for len(queue) != 0 {
		level := make([]int,0)
		//申请一个变量记录长度
		size := len(queue)
		for i := 0;i < size;i++ {
			//弹出队列头节点
			temp := queue[0]
			queue = queue[1:]
			//将队列头节点的左右子节点入队 若有
			if temp.Left != nil {
				queue = append(queue,temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue,temp.Right)
			}
			level = append(level,temp.Val)
		}
		if reverse {
			level = reverseSlice(level)
		}
		reverse = !reverse
		result = append(result,level)

	}
	return result

}

func reverseSlice(nums []int) []int {
	reverse := make([]int,len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		reverse[len(nums)-1-i] = nums[i]
	}
	return reverse
}
