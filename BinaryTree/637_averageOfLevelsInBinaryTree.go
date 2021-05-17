/**
 leetcode 637 Average of Levels in Binary Tree

 Given the root of a binary tree, return the average value of the nodes on each
 level in the form of an array. Answers within 10-5 of the actual answer will be accepted.

 */
package main

import (
	"fmt"
	"strconv"
)


func averageOfLevels(root *TreeNode) []float64 {
	queue := make([]*TreeNode,0)
	var result []float64
	if root == nil {
		return nil
	}
	//入队 root节点 第一层
	queue = append(queue,root)
	for len(queue) != 0 {
		sum := 0
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
			sum += temp.Val
		}
		value, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", float64(sum)/float64(size)), 64)
		result = append(result,value)

	}
	return result

}
