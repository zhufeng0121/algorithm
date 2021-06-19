/**
 leetcode 121 Best Time to Buy and Sell Stock

 You are given an array prices where prices[i] is the price
 of a given stock on the ith day.

 You want to maximize your profit by choosing a single day
 to buy one stock and choosing a different day in the future
 to sell that stock.

 Return the maximum profit you can achieve from this transaction.
 If you cannot achieve any profit, return 0

 */
package main
//买卖股票的最佳时机

// 1.思路1： 计算数组前后相邻元素的差值，然后计算最大连续子数组和
func maxProfit(prices []int) int {
	pricesII := make([]int,0)
	for i := 1;i < len(prices);i++ {
		pricesII = append(pricesII, prices[i] - prices[i-1])
	}

	result := maxSubarraySum(pricesII)
	if result <= 0 {
		return 0
	}
	return result

}

func maxSubarraySum(prices []int) int {
	max := ^INT_MAX
	cur := 0
	for i := 0;i < len(prices);i++ {
		cur += prices[i]
		if cur < 0 {
			cur = 0
		} else {
			if max < cur {
				max = cur
			}
		}
	}
	return max
}

//dp[i][k][0 or 1]
//0 <= i <= n-1, 1 <= k <= K
//n 为天数，大 K 为最多交易数
//此问题共 n × K × 2 种状态，全部穷举就能搞定


//这个题很经典，后面还有许多有意思的解法，需要继续探究学习
//dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
//dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0] - prices[i])
//= max(dp[i-1][1][1], -prices[i])
//解释：k = 0 的 base case，所以 dp[i-1][0][0] = 0

//dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
//max(   选择 rest  ,             选择 sell      )
//
//解释：今天我没有持有股票，有两种可能：
//要么是我昨天就没有持有，然后今天选择 rest，所以我今天还是没有持有；
//要么是我昨天持有股票，但是今天我 sell 了，所以我今天没有持有股票了。
//
//dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
//max(   选择 rest  ,           选择 buy         )
//
//解释：今天我持有着股票，有两种可能：
//要么我昨天就持有着股票，然后今天选择 rest，所以我今天还持有着股票；
//要么我昨天本没有持有，但今天我选择 buy，所以今天我就持有股票了。



func maxProfitDP(prices []int) int {
	n := len(prices)
	//声明一个 n * 2的动态规划数组
	dp := make([][]int,n)
	for i := 0; i < n;i++ {
		dp[i] = make([]int,2)
	}
	for i := 0; i < n;i++ {
		//base case
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0],dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])

	}
	return dp[n-1][0]

}




