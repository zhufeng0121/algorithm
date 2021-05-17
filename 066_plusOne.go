/**
 leetcode 66 Plus One

 Given a non-empty array of decimal digits representing a non-negative
 integer, increment one to the integer.The digits are stored such that
 the most significant digit is at the head of the list, and each element
 in the array contains a single digit.

 You may assume the integer does not contain any leading zero, except
 the number 0 itself.
 */
package main

func plusOne(digits []int ) []int {
	if len(digits) == 0 {
		return []int{}
	}
	carry := 1
	//注意边界
	for i := len(digits) -1; i >= 0 ;i-- {
		if digits[i] + carry > 9 {
			digits[i] = 0
			carry = 1
		} else {
			digits[i] += carry
			carry = 0
		}
	}
	//向切片中头部插入元素，这个写法主意一下
	if digits[0] == 0 && carry == 1 {
		digits = append([]int{1},digits...)
	}
	return digits

}


