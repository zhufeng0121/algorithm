/**
 leetcode 714  Best Time to Buy and Sell Stock with Transaction Fee

 You are given an array prices where prices[i] is the price
 of a given stock on the ith day, and an integer fee representing
 a transaction fee.

 Find the maximum profit you can achieve. You may complete as
 many transactions as you like, but you need to pay the
 transaction fee for each transaction.

 Note: You may not engage in multiple transactions simultaneously
 (i.e., you must sell the stock before you buy again).
 */
package main

//dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i] - fee)
//dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
func maxProfitWithTransaction(prices []int, fee int) int {
	n := len(prices)
	//声明一个 n * 2的动态规划数组
	dp := make([][]int,n)
	for i := 0; i < n;i++ {
		dp[i] = make([]int,2)
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[0]- fee
			continue
		}

		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i] - fee)

	}
	return dp[n-1][0]

}
