/**
 leetcode 590. N-ary Tree Postorder Traversal

 Given the root of an n-ary tree, return the postorder traversal
 of its nodes' values.

 Nary-Tree input serialization is represented in their level order
 traversal. Each group of children is separated by the null value
 (See examples)

 */
package main

//N-ary 后续遍历，左，右，中
func postorder(root *NaryNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int,0)
	for i := 0; i < len(root.Children);i++ {
		res = append(res, postorder(root.Children[i])...)
	}
	res = append(res, root.Val)
	return res

}

//iterative 这种写法快好多呀
func postorderI(root *NaryNode) []int {
	T := []*NaryNode{}
	if root != nil {
		T = append(T, root)
	}
	rt := []int{}

	for {
		if len(T) == 0 {
			break
		}

		node := T[len(T)-1]
		T = T[:len(T)-1]
		T = append(T, node.Children...)
		rt = append(rt, node.Val)
	}

	for i, j := 0, len(rt)-1; i < j; i, j = i+1, j-1 {
		rt[i], rt[j] = rt[j], rt[i]
	}

	return rt
}
