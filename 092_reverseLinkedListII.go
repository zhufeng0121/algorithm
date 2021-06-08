/**
 leetcode 92 Reverse Linked List II

 Given the head of a singly linked list and two integers left and right where
 left <= right, reverse the nodes of the list from position left to position right,
 and return the reversed list.

*/
package main

//思路是对的，需要保存left前的一个节点，
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	//题目中给出假设, 不需要对left和right的值进行校验，直接进行使用即可，left 取值为1代表的是头节点
	//需要找到left节点的前一个节点，和right节点的后一个节点
	virHead := &ListNode{
		0,head,
	}
	tmp := virHead
	//left 节点的前一个节点
	var before *ListNode
	var leftNode * ListNode
	//right 节点的后一个节点
	var behind *ListNode
	var rightNode *ListNode
	index := 0
	for tmp != nil {
		if index == left - 1  {
			before = tmp
			leftNode = tmp.Next
		} else if index == right{
			rightNode = tmp
			behind = tmp.Next
		}
		index++
		tmp = tmp.Next
	}
	// 切出一个子链表
	before.Next = nil
	rightNode.Next = nil

	//对子链表进行翻转
	myReverse(leftNode)

	before.Next = rightNode

	leftNode.Next = behind

	return virHead.Next

}
//迭代法
func myReverse(head *ListNode) {
	var prev *ListNode
	curr := head
	for curr != nil {
		//next 记录当前节点的Next节点
		next := curr.Next
		//prev赋值 curr.Next
		curr.Next = prev
		//将prev 赋值为 curr
		prev = curr
		// curr 赋值为 next
		curr = next
	}
}

