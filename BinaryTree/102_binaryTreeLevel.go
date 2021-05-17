/**
 leetcode 102 Binary Tree Level Order Traversal

 Given the root of a binary tree, return the level order traversal of its nodes'
 values. (i.e., from left to right, level by level).

 */
package main

// 层次遍历 需要队列
// 我们可以利用队列先进先出的特性，使用队列来帮助我们完成层序遍历, 具体操作如下
//
//让二叉树的每一层入队, 然后再依次执行出队操作,
//
//对该层节点执行出队操作时, 需要将该节点的左孩子节点和右孩子节点进行入队操作,
//
//这样当该层的所有节点出队结束后, 下一层也就入队完毕,
//
//不过我们需要考虑的就是, 我们需要通过一个变量来保存每一层节点的数量.
func levelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode,0)
	var result [][]int
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
	return result

}
