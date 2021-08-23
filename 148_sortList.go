/**
 leetcode 148 Sort List

 Given the head of a linked list, return the list after sorting it in ascending order.

 Follow up: Can you sort the linked list in O(n logn) time and O(1) memory
 (i.e. constant space)?


 */
package main

//合并排序
func sortList(head *ListNode) *ListNode {
	//递归结束条件
	if head == nil || head.Next == nil {
		return head
	}

	//2.找到链表中间节点并断开链表
	midNode := GetMiddleNode(head)
	rightNode := midNode.Next
	midNode.Next = nil

	left := sortList(head)
	right := sortList(rightNode)

	return MergeTwoSortedList(left,right)

}

//利用快慢指针获取链表的中间节点
func GetMiddleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//此处fast的初始值如果为head,那么当链表中元素的个数为偶数时，返回的是两个中间节点的后一个节点
	//如果初始值为head.Next.Next，则返回的是中间节点的前一个节点
	slow ,fast := head, head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
//合并两个有序链表
func MergeTwoSortedList(l1 *ListNode, l2 *ListNode) *ListNode {
	vir := &ListNode{Val: -1}
	cur := vir
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}

	return vir.Next
}