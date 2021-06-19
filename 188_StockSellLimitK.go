/**
 leetcode 188 Best Time to Buy and Sell Stock IV

 You are given an integer array prices where prices[i] is
 the price of a given stock on the ith day, and an integer k.

 Find the maximum profit you can achieve. You may complete at
 most k transactions.

 Note: You may not engage in multiple transactions simultaneously
 (i.e., you must sell the stock before you buy again).

 */
package main

func maxProfitLimitK(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][][]int,n)
	for i := 0; i < n;i++ {
		dp[i] = make([][]int,k + 1)
		for j := 0; j < k +1 ;j++ {
			dp[i][j] = make([]int,2)
		}
	}
	for i := 0 ;i < n; i++ {
		for j := k; j >= 1; j-- {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[0]
				continue

			}
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1] + prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0] - prices[i])
		}
	}

	return dp[n-1][k][0]

}