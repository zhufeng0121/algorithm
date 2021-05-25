/**
 leetcode 538 Convert BST to Greater Tree

 Given the root of a Binary Search Tree (BST), convert it to a Greater Tree
 such that every key of the original BST is changed to the original key plus
 sum of all keys greater than the original key in BST.

 As a reminder, a binary search tree is a tree that satisfies these constraints:

 The left subtree of a node contains only nodes with keys less than the node's key.
 The right subtree of a node contains only nodes with keys greater than the node's key.
 Both the left and right subtrees must also be binary search trees.


 */
//这个题的动机是我们应该先修改数值比较大的节点，然后修改数值比较小的节点。
//这样做，才能保证，我们只需要一次遍历，就把每个节点的值修改成了比它值更大的所有节点的和。

//BST的右子树都比该节点大，所以修改次序是右–>中–>左。用一个变量储存遍历过程中所有有节
//点之和就得到了所有的比当前把该节点的值更大的和，然后修改为该变量的值即可。

package BinarySearchTree

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	helper(root, &sum)
	return root
}

func helper(root *TreeNode, sum *int) {
	if root == nil { return }
	helper(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	helper(root.Left, sum)
}

