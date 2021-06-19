/**
 leetcode 24 Swap Nodes in Pairs

 Given a linked list, swap every two adjacent nodes and return its head.
 You must solve the problem without modifying the values in the list's
 nodes (i.e., only nodes themselves may be changed.)

 */
package main

// 利用递归进行两两交换
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//head.Next是新的头节点
	newHead := head.Next
	//head被交换到第二个节点，head.Next 是 原第二个节点的newHead的Next
	head.Next = swapPairs(newHead.Next)
	//新头结点的Next 为head
	newHead.Next = head
	return newHead
}
//创建虚拟头结点 交换之前的节点关系是 temp -> node1 -> node2，交换之后的节点关系要变成 temp -> node2 -> node1
func swapPairsI(head *ListNode) *ListNode  {
	virHead := &ListNode{Val: 0,Next: head}
	tmp := virHead
	//这个条件yao保证tmp1后面有两个节点
	for tmp.Next != nil && tmp.Next.Next != nil {
		//获取tmp1的后二节点
		node1 := tmp.Next
		node2 := tmp.Next.Next
		tmp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		//node1 已经变成了tmp.Next.Next
		tmp = node1
	}
	return virHead.Next

}

//这题有一种变形操作，这是交换两个一组的节点，如果是k个一组的节点，可单独封装一个函数，将k作为参数传入
