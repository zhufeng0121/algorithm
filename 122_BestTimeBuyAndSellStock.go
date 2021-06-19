/**
 leetcode 122 Best Time to Buy and Sell Stock

 You are given an array prices where prices[i] is the price
of a given stock on the ith day.

 Find the maximum profit you can achieve. You may complete
 as many transactions as you like (i.e., buy one and sell
 one share of the stock multiple times).

 Note: You may not engage in multiple transactions simultaneously
 (i.e., you must sell the stock before you buy again).


 */
package main
//可以卖出多次，求最大利润
//见121题解
//此时交易次数没有限制
//如果 k 为正无穷，那么就可以认为 k 和 k - 1 是一样的。可以这样改写框架：
//dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
//dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
//= max(dp[i-1][k][1], dp[i-1][k][0] - prices[i])
//我们发现数组中的 k 已经不会改变了，也就是说不需要记录 k 这个状态了：


//dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
//dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
func maxProfitII(prices []int) int {
	n := len(prices)
	dp := make([][]int,n)
	for i := 0; i < n;i++ {
		dp[i] = make([]int,2)
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0],dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][1],dp[i-1][0] - prices[i])

	}
	return dp[n-1][0]

}
