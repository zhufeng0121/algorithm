/**
 leetcode 172 Factorial Trailing Zeroes

 Given an integer n, return the number of trailing zeroes in n!

 */
package main

//只有 2 和 5的乘积才会出现0 ，统计 5出现的次数，2出现的次数，2一定出现的次数比5多，所以仅统计5出现的次数就行
func trailingZeroes(n int) int {
	count := 0
	for i := 5; i <= n;i +=5 {
		cur := i
		for cur >= 5 {
			if cur % 5 == 0 {
				count++
				cur = cur / 5
			} else {
				break
			}
		}
	}
	return count
}

func trailingZeroes2(n int) int {
	if n / 5 == 0 {
		return 0
	}
	return n/5 + trailingZeroes2(n/5)
}


