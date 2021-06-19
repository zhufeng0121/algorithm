/**
 leetcode 99 Recover Binary Search Tree

 You are given the root of a binary search tree (BST), where
 exactly two nodes of the tree were swapped by mistake.
 Recover the tree without changing its structure.

 Follow up: A solution using O(n) space is pretty straight forward.
 Could you devise a constant space solution?
 */
package main

//中序遍历： BST的中序遍历是有序的数组，如果中序遍历的数组中出现两次下降，那么
//第一个错误的数字是第一次下降的较大的数，第二个错误的数字是第二次下降的较小的数
//如果只有一次下降，那么下降较大的数和较小的数就是两个错误的节点，交换即可

//层次遍历是广度优先搜索BFS(using queue) ,此处是中序遍历，DFS，using stack

func recoverTree(root *TreeNode)  {
	if root == nil {
		return
	}
	var node1 *TreeNode
	var node2 *TreeNode
	//记录遍历节点的前序节点，要将当前节点和前序节点进行比较
	var pre   *TreeNode
	stack := make([]*TreeNode,0)
	for len(stack) != 0 || root != nil {
		if root != nil {
			stack = append(stack,root)
			root = root.Left
		} else {
			root = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			// 这里还是存疑
			if pre != nil && pre.Val > root.Val {
				if node1 == nil {
					node1 = pre
				}
				node2 = root
			}
			pre = root
			root = root.Right
		}
	}
	if node1 != nil && node2 != nil {
		node1.Val, node2.Val = node2.Val,node1.Val

	}

}
