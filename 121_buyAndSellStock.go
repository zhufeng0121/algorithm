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


//这个题很经典，后面还有许多有意思的解法，需要继续探究学习






