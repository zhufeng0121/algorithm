/**
 leetcode 236 Lowest Common Ancestor of a Binary Tree

 Given a binary tree, find the lowest common ancestor
 (LCA) of two given nodes in the tree.

 According to the definition of LCA on Wikipedia: “The lowest common
 ancestor is defined between two nodes p and q as the lowest node in
 T that has both p and q as descendants (where we allow a node to be
 a descendant of itself).”


 */
package main

func lowestCommonAncestorIII(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestorIII(root.Left,p,q)
	right := lowestCommonAncestorIII(root.Right,p,q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right

}

func lowestCommonAncestorIV(root, p, q *TreeNode) *TreeNode {
	pathP := make([]*TreeNode,0)
	pathQ := make([]*TreeNode,0)
	var ancestor *TreeNode
	flag1, flag2 := searchNode(root,pathP,p), searchNode(root,pathP,q)
	if flag1 && flag2 {
		for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
			ancestor = pathP[i]
		}

	}
	return ancestor

}

func searchNode(root *TreeNode,path []*TreeNode, target *TreeNode) bool {
	if root == nil {
		return false
	}
	path = append(path,root)
	if root.Val == target.Val {
		return true
	}
	flag := false
	//先去左子树找
	if root.Left != nil {
		flag = searchNode(root.Left,path,target)
	}
	//左子树找不到并且右子树不为空的情况下才去找
	if !flag && root.Right != nil {
		flag = searchNode(root.Right,path,target)
	}
	//左右都找不到，弹出栈顶元素
	if !flag {
		path = path[0:len(path) -1]
	}
	return flag

}