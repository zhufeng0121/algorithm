/**
 leetcode 67 Add Binary

 Given two binary strings a and b,
 return their sum as a binary string.

 */

package main

import "strconv"

// 思路是在短的数组前面不断补0,使长度和长数组相同，然后进行逐位相加
func addBinary(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}
	res := make([]string,len(a) + 1)
	i, j, k, c := len(a)-1 ,len(b)-1,len(a),0
	for i >= 0 && j >=0 {

	}
}
