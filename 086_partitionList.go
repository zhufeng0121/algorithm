/**
 leetcode 86 Partition List

 Given the head of a linked list and a value x, partition it such
 that all nodes less than x come before nodes greater than or equal to x.

 You should preserve the original relative order of the nodes in
 each of the two partitions.

*/
package main

//官方题解中给出的是 维护一个大于和小于的两个链表，实际上比较类似
//只不过自己的解法是维护一个大于和小于的数组，逻辑上差不多
func partition(head *ListNode, x int) *ListNode {
	if head == nil  {
		return nil
	} else if head.Next == nil {
		return head
	}
	less := make([]int,0)
	more := make([]int,0)
	cur := head
	for cur != nil {
		if cur.val < x {
			less = append(less,cur.val)
		} else {
			more = append(more,cur.val)
		}
		cur = cur.Next
	}
	newHead := &ListNode{val: 0,Next: nil}
	cur = newHead
	for _, v := range less {
		cur.Next = &ListNode{val: v,Next: nil}
		cur = cur.Next
	}
	for _, v := range more {
		cur.Next = &ListNode{val: v,Next: nil}
		cur = cur.Next
	}
	return newHead.Next


}
//这个题还有其他的做法，比如双指针，研究一下
func partitionI(head *ListNode, x int) *ListNode {

}


