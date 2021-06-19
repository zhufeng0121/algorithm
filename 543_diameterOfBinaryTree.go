/**
 leetcode 543. Diameter of Binary Tree

 Given the root of a binary tree, return the length
 of the diameter of the tree.

 The diameter of a binary tree is the length of the longest
 path between any two nodes in a tree. This path may or may
 not pass through the root.

 The length of a path between two nodes is represented by the
 number of edges between them.

*/
package main
//如果 root 存左子树和右子数的话，那么diameter 必然经过root，否则，
func diameterOfBinaryTree(root *TreeNode) int {
	return helperV(root).maxDistance

}

type returnType struct {
	maxDistance int
	height      int
}

func helperV(root *TreeNode) *returnType {
	if root == nil {
		return &returnType{0,0}
	}
	left  := helperV(root.Left)
	right := helperV(root.Right)
	height := max(left.height,right.height) + 1
	maxDistance := max(max(left.maxDistance,right.maxDistance), left.height + right.height)
	return &returnType{height: height,maxDistance: maxDistance}

}


//试试第二种解法,不过第二种解法是真的慢的说，用了太多次递归了
//思路还是
func getMaxDistance(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getMaxDistance(root.Right),getMaxDistance(root.Left)) + 1
}
func diameterOfBinaryTreeI(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cross := getMaxDistance(root.Left)+getMaxDistance(root.Right)
	return max(max(cross,diameterOfBinaryTreeI(root.Left)),diameterOfBinaryTreeI(root.Right))

}