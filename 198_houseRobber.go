/**
 leetcode 198 House Robber

 You are a professional robber planning to rob houses along a street.
 Each house has a certain amount of money stashed, the only constraint
 stopping you from robbing each of them is that adjacent houses have
 security systems connected and it will automatically contact the police
 if two adjacent houses were broken into on the same night.

 Given an integer array nums representing the amount of money of each house,
 return the maximum amount of money you can rob tonight without alerting the police.

 */
package main
//DP
// 设 profit[i] 为抢劫前i户人家所能获得的最大钱数 那么 profit[i] = max(profit[i-1],profit[i-2] + num[i])
func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	profit := make([]int, len(nums))
	profit[0] = nums[0]
	profit[1] = maxInt(nums[0],nums[1])
	result := maxInt(profit[0],profit[1])
	for i := 2 ;i < len(nums);i++ {
		profit[i] = maxInt(profit[i-1], profit[i-2] + nums[i])
		result = maxInt(result,profit[i])
	}
	return result

}
