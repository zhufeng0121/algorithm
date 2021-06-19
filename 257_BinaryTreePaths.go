/**
 leetcode 257 Binary Tree paths

 Given the root of a binary tree, return all root-to-leaf paths in any order.

 A leaf is a node with no children.

*/
package main

import (
	"fmt"
	"strconv"

)

//DFS
func binaryTreePaths(root *TreeNode) []string {
	res :=[]string{}
	binaryTreePathsHelper(root,&res,"")
	return res

}

func binaryTreePathsHelper(node *TreeNode,res *[]string, temp string) {
	if node == nil {
		return
	}
	if len(temp) == 0 {
		temp = strconv.Itoa(node.Val)
	} else {
		temp += "->" + strconv.Itoa(node.Val)
	}
	if node.Left == nil && node.Right == nil {
		*res = append(*res,temp)
		return
	}
	binaryTreePathsHelper(node.Left, res, temp)
	binaryTreePathsHelper(node.Right, res, temp)

}
func binaryTreePathsI(root *TreeNode) []string {
	paths := make([]string, 0)
	if root != nil {
		DFS(root, "", &paths)
	}
	return paths
}

func DFS(node *TreeNode, p string, paths *[]string) {
	if node.Left == nil && node.Right == nil {
		*paths = append(*paths, fmt.Sprintf("%s%d", p, node.Val))
	} else {
		path := fmt.Sprintf("%s%d->", p, node.Val)
		if node.Left != nil {
			DFS(node.Left, path, paths)
		}
		if node.Right != nil {
			DFS(node.Right, path, paths)
		}
	}
}


func binaryTreePathsII(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%d", root.Val)}
	}

	l := binaryTreePaths(root.Left)
	r := binaryTreePaths(root.Right)

	for i := 0; i < len(l); i++ {
		l[i] = fmt.Sprintf("%d->%s", root.Val, l[i])
	}

	for i := 0; i < len(r); i++ {
		r[i] = fmt.Sprintf("%d->%s", root.Val, r[i])
	}

	res := append(l, r...)
	return res
}
