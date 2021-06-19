/**
 leetcode 123  Best Time to Buy and Sell Stock III

 You are given an array prices where prices[i] is the price
 of a given stock on the ith day.

 Find the maximum profit you can achieve. You may complete
 at most two transactions.

 Note: You may not engage in multiple transactions simultaneously
 (i.e., you must sell the stock before you buy again).

*/
package main

//见122题解
//此时交易次数 限制为2次，因此 k = 2

//dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
//dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])

func maxProfitLimitV(prices []int) int {
	k := 2
	n := len(prices)
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
//最多交易两次
func maxProfitLimit(prices []int) int {
	n := len(prices)
	dp := make([][][]int,n)
	for i := 0; i < n;i++ {
		dp[i] = make([][]int,3)
		for j := 0; j < 3;j++ {
			dp[i][j] = make([]int,2)
		}
	}
	dp[0][0][0] = 0
	dp[0][1][0] = 0
	dp[0][1][1] = -prices[0]
	dp[0][2][0] = 0
	dp[0][2][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][2][0] = max(dp[i-1][2][0],dp[i-1][2][1] + prices[i])
		dp[i][2][1] = max(dp[i-1][2][1],dp[i-1][1][0] - prices[i])
		dp[i][1][0] = max(dp[i-1][1][0],dp[i-1][1][1] + prices[i])
		dp[i][1][1] = max(dp[i-1][1][1],dp[i-1][0][0] - prices[i])
	}
	return dp[n-1][2][0]
}
