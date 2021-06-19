/**
 leetcode 235 Lowest Common Ancestor of a Binary Search Tree

 Given a binary search tree (BST), find the lowest common ancestor
 (LCA) of two given nodes in the BST.

 According to the definition of LCA on Wikipedia: “The lowest common
 ancestor is defined between two nodes p and q as the lowest node in
 T that has both p and q as descendants (where we allow a node to be
 a descendant of itself).”
 */
package main

//如果p和q分列root节点的左右两侧，说明根结点就是LCA，否则，进行递归即可
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	return findLCA(root,p,q)

}
func findLCA(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val <= root.Val && root.Val <= q.Val {
		return root
	}
	if p.Val > root.Val {
		return findLCA(root.Right,p,q)
	}
	if q.Val < root.Val {
		return findLCA(root.Left,p,q)
	}

	return nil
}

func getPath(root, target *TreeNode) (path []*TreeNode) {
	node := root
	for node != target {
		path = append(path, node)
		if target.Val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	path = append(path, node)
	return
}

func lowestCommonAncestorII(root, p, q *TreeNode) (ancestor *TreeNode) {
	pathP := getPath(root, p)
	pathQ := getPath(root, q)
	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
		ancestor = pathP[i]
	}
	return
}

