/**
 leetcode 343 Integer Break

 Given an integer n, break it into the sum of k positive integers,
 where k >= 2, and maximize the product of those integers.

 Return the maximum product you can get.
 */
package main


func intergerBreak(n int) int {
	if n < 2 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		mx := 0
		for j :=1; j < i;j++ {
			//这个含义没有看懂
			mx = maxInt(mx, maxInt(dp[j],j) *maxInt(i-j,dp[i-j]))
		}
		dp[i] = mx
	}
	return dp[n]
}


//DP : 状态转移方程 dp[i] = max()
