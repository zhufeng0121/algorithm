/**
 leetcode 24 Swap Nodes in Pairs

 Given a linked list, swap every two adjacent nodes and return its head.
 You must solve the problem without modifying the values in the list's
 nodes (i.e., only nodes themselves may be changed.)

 */
package main

//看来我这么想还真没错
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur, temp := head, head
	for cur != nil && cur.Next != nil {
		temp = cur.Next
		cur.Next = cur.Next.Next
		cur = temp
		cur = cur.Next.Next

	}

}
// 可以利用数组结构 （也可以利用栈结构，每K个为一组）
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	arr := make([]int,0)
	cur := head
	for cur.Next != nil {
		arr = append(arr, cur.val)
	}
	n := len(arr)
	for i := 0 ;i < n;i += 2 {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}
	cur = head
	n--
	for n != 0 {
		cur.Next = &ListNode{val: arr[n],Next: nil}
		n--
	}
	return cur
}


