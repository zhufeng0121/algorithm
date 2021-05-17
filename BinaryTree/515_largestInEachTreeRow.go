/**
 leetcode 515 Find Largest Value in Each Tree Row

 Given the root of a binary tree, return an array of the largest value in
 each row of the tree (0-indexed).

 */
package main
const MAX = int(^uint32(0) >> 1)
func largestValues(root *TreeNode) []int {
	queue := make([]*TreeNode,0)
	var result []int
	if root == nil {
		return nil
	}
	//入队 root节点 第一层
	queue = append(queue,root)
	for len(queue) != 0 {
		max := ^MAX
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
			max = maxInt(max,temp.Val)

		}
		result = append(result,max)

	}
	return result

}
func maxInt(a,b int) int {
	if a > b {
		return a
	}
	return b
}