/**
 leetcode 002 add two number

 */
package main

type ListNode2 struct {
	Val int
    Next *ListNode2
}


func addTwoNumbers(l1 *ListNode2,l2 *ListNode2) *ListNode2 {
	if l1 == nil  || l2 == nil {
		return nil
	}
	//增加虚拟头结点
	head := &ListNode2{Val:0,Next:nil}
	current := head
	carry := 0
	for l1 != nil || l2 != nil {
		var x,y int
		if l1 == nil {
			x = 0
		}else {
			x = l1.Val
		}
		if l2 == nil {
			y = 0
		}else {
			y = l2.Val
		}
		current.Next = &ListNode2{Val:(x+y+carry)%10,Next:nil}
		current = current.Next
		carry = (x+y+carry)/10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		current.Next = &ListNode2{Val: carry %10, Next:nil}
	}

	return head.Next

}