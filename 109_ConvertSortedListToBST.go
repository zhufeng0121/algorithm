/**
 leetcode 109 Convert Sorted List to Binary Search Tree

 Given the head of a singly linked list where elements are
 sorted in ascending order, convert it to a height balanced BST.

 For this problem, a height-balanced binary tree is defined as
 a binary tree in which the depth of the two subtrees of every
 node never differ by more than 1.

 */
package main

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}
	mid := getMiddleNodeI(head,tail)
	next := mid.Next
	root := &TreeNode{Val: mid.Val}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(next)
	return root
}

//寻找中间节点 可以采用双指针
func getMiddleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow

}

func sortedListToBSTI(head *ListNode) *TreeNode {
	return helperV(head,nil)
}

func helperV(head *ListNode,end *ListNode) *TreeNode {
	if head == end {
		return nil
	}
	if head.Next == end {
		return &TreeNode{Val: head.Val}
	}
	mid := getMiddleNodeI(head,end)
	return &TreeNode{
		Val: mid.Val,
		Left: helperV(head,mid),
		Right:helperV(mid.Next,end),
	}
}

//获取中间节点
func getMiddleNodeI(head *ListNode,end *ListNode) *ListNode {
	fast, slow := head, head
	for fast != end && fast.Next != end {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow

}

//将有序链表转换成有序数组
func sortedListToBSTII(head *ListNode) *TreeNode {
	arr := []int{}
	cur := head
	for cur != nil {
		arr = append(arr,cur.Val)
		cur = cur.Next
	}
	return buildBST(arr,0,len(arr) -1)

}
func buildBST(arr []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end - start) >> 1
	root := &TreeNode{
		Val: arr[mid],
	}
	root.Left = buildBST(arr,start,mid -1)
	root.Right = buildBST(arr,mid+1,end)
	return root

}