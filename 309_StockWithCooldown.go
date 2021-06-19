/**
 leetcode 309 Best Time to Buy and Sell Stock with Cooldown

 You are given an array prices where prices[i] is the
 price of a given stock on the ith day.

 Find the maximum profit you can achieve. You may complete
 as many transactions as you like (i.e., buy one and sell
 one share of the stock multiple times) with the following
 restrictions:

 After you sell your stock, you cannot buy stock on the next
 day (i.e., cooldown one day).
 Note: You may not engage in multiple transactions simultaneously
 (i.e., you must sell the stock before you buy again).
 */

package main


//dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
//dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
func maxProfitWithCool(prices []int) int {
	n := len(prices)

	if n < 2 {
		return 0
	}
	if n == 2 {
		return max(prices[1] - prices[0], 0)
	}
	dp := make([][]int,n)
	for i := 0; i < n;i++ {
		dp[i] = make([]int,2)
	}
	dp[0][1] = -prices[0]
	dp[1][0] = max(dp[0][0],dp[0][1] + prices[1])
	dp[1][1] = max(dp[0][1],dp[0][0] - prices[1])
	for i := 2; i < n; i++ {
		dp[i][0] = max(dp[i-1][0],dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][1],dp[i-2][0] - prices[i])

	}
	return dp[n-1][0]

}