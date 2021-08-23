package main

func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	//将l1 和 l2
	head1, head2 := l1, l2
	stack1 := make([]int,0)
	stack2 := make([]int,0)
	for head1 != nil {
		stack1 = append(stack1, head1.Val)
		head1 = head1.Next
	}
	for head2 != nil {
		stack2 = append(stack2,head2.Val)
		head2 = head2.Next
	}
	var node *ListNode
	carry := 0
	//如果
	for len(stack1) != 0 || len(stack2) != 0 || carry > 0 {
		sum := carry
		if len(stack1) != 0 {
			sum += stack1[len(stack1) - 1]
			stack1 = stack1[:len(stack1) - 1]
		}
		if len(stack2) != 0 {
			sum += stack2[len(stack2) - 1]
			stack2 = stack2[:len(stack2) - 1]
		}
		tmp := &ListNode{Val: sum %10}
		tmp.Next = node
		node = tmp
		carry = sum / 10
	}
	return node

}


