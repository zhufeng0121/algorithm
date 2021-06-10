/**
 leetcode 342 Power of 4

 Given an integer n, return true if it is a power of two.

 Follow up: Could you solve it without loops/recursion?


 */
package main

//如果 n 是 4 的幂，那么 n 的二进制表示中有且仅有一个 1，并且这个 1
//出现在从低位开始的第偶数个二进制位上
//
//由于题目保证了 nn 是一个 32位的有符号整数，因此我们可以构造一个整数
//mask，它的所有偶数二进制位都是 0，所有奇数二进制位都是 1。 这样一来，
//我们将 n 和mask 进行按位与运算， 如果结果为 0，说明 n 二进制表示中的
//1 出现在偶数的位置，否则说明其出现在奇数的位置。
//
//mask 的二进制表示为:
//mask=(10101010101010101010101010101010)
//mask=(AAAAAAAA)

func isPowerOfFour(n int) bool {
	return n > 0 && n & (n-1) == 0 && n & 0xAAAAAAAA == 0
}

