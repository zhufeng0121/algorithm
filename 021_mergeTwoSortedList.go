/**
 leetcode 021 Merge Two Sorted Lists

 Merge two sorted linked lists and return it as a sorted list.
 The list should be made by splicing together the nodes of the first two lists.

 */
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{Val: 0,Next: nil}
	current := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = &ListNode{Val: l1.Val,Next: nil}
			l1 = l1.Next
			current = current.Next
		} else {
			current.Next = &ListNode{Val: l2.Val,Next: nil}
			l2 = l2.Next
			current = current.Next
		}

		if l1 == nil {
			for l2 != nil {
				current.Next = &ListNode{Val: l2.Val,Next: nil}
				l2 = l2.Next
				current = current.Next
			}
		}
		if l2 == nil {
			for l1 != nil {
				current.Next = &ListNode{Val: l1.Val,Next: nil}
				l1 = l1.Next
				current = current.Next
			}
		}
	}
	return head.Next
}

//迭代合并两个有序链表
func MergeTwoSortedListI(l1 *ListNode, l2 *ListNode) *ListNode {
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
func mergeTwoListsII(l1 *ListNode, l2 *ListNode) *ListNode {
	//TODO: 可以用递归的写法再试一下
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	newHead := &ListNode{}
	if l1.Val < l2.Val {
		newHead = l1
		newHead.Next = mergeTwoListsII(l1.Next,l2)
	} else {
		newHead = l2
		newHead.Next = mergeTwoListsII(l1,l2.Next)
	}
	return newHead
}



