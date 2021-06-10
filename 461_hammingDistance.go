/**
 leetcode 461 Hamming Distance

 The Hamming distance between two integers is the number of positions
 at which the corresponding bits are different.

 Given two integers x and y, return the Hamming distance between them.

*/
package main
func hammingDistance(x int, y int) int {
	xor := x ^ y
	count := 0
	for xor != 0 {
		// 注意 n 和 n-1 之间 &运算加括号
		xor = xor & (xor-1)
		count++

	}
	return count
}


