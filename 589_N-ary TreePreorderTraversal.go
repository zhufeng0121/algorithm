/**
 leetcode 589 N-ary Tree Preorder Traversal

 Given the root of an n-ary tree, return the preorder
 traversal of its nodes' values.

 Nary-Tree input serialization is represented in their
 level order traversal. Each group of children is separated
 by the null value (See examples)
 */
package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type NaryNode struct {
	Val int
	Children []*NaryNode
}

func preorder(root *NaryNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int,0)
	res = append(res, root.Val)
	for i := 0; i < len(root.Children); i++ {
		res = append(res, preorder(root.Children[i])...)
	}
	return res

}
