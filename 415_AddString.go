/**
 leetcode 415 Add String

 Given two non-negative integers, num1 and num2 represented as
 string, return the sum of num1 and num2 as a string.

 You must solve the problem without using any built-in library
 for handling large integers (such as BigInteger). You must
 also not convert the inputs to integers directly.

 */
package main

import "strings"

//Input: num1 = "11", num2 = "123"
//Output: "134"

//
func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	num2 = strings.Repeat("0",len(num1) - len(num2)) + num2
	arr1 := []byte(num1)
	arr2 := []byte(num2)
	carry := byte(0)
	res := make([]byte,len(num1))
	for i := len(num1) - 1 ; i >= 0; i-- {
		res[i] = (carry + arr1[i] - '0' + arr2[i] - '0') % 10 +'0'
		carry  = (carry + arr1[i] - '0' + arr2[i] - '0') / 10
	}

	if carry == 1 {
		return "1" + string(res)
	}

	return string(res)

}
