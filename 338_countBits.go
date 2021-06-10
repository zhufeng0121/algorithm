/**
 leetcode 338. Counting Bits

 Given an integer n, return an array ans of length n + 1 such that for each i
 (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

*/
package main

func countBits(n int) int {
	res := make([]int, n + 1)
	for i := 0;i <= n;i++ {
		res[i] = numberOfOne(i)
	}
	return res

}
func numberOfOne(n int) int {
	count := 0
	for n != 0 {
		// 注意 n 和 n-1 之间 &运算加括号
		n = n & (n-1)
		count++

	}
	return count
}

//利用奇偶性来进行判断
//1.奇数：二进制表示中，奇数一定比前面那个偶数多一个 1，因为多的就是最低位的1
//2.偶数：二进制表示中，偶数中 1 的个数一定和除以 2 之后的那个数一样多。因为最低位是 0，
//除以 2 就是右移一位，也就是把那个 0 抹掉而已，所以 1 的个数是不变的。

func numberOfOneII(n int) []int {
	res := make([]int, n + 1)
	res[0] = 0
	for i := 1;i <= n; i++ {
		if i & 1 == 1 {
			res[i] = res[i-1] + 1
		} else {
			res[i] = res[i/2]
		}
	}
	return res
}
//DP
// 递推公式 answer[i] = answer[i >> 1] + (i & 1)

func countBitsIV(n int) []int {
	res := make([]int, n + 1)
	res[0] = 0
	for i := 1;i <= n; i++ {
		res[i] = res[i >> 1] + (i & 1)

	}
	return res
}
