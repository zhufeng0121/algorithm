/**
 leetcode 213. House Robber II

 You are a professional robber planning to rob houses along a street.
 Each house has a certain amount of money stashed. All houses at this
 place are arranged in a circle. That means the first house is the
 neighbor of the last one. Meanwhile, adjacent houses have a security
 system connected, and it will automatically contact the police if two
 adjacent houses were broken into on the same night.

 Given an integer array nums representing the amount of money of each
 house, return the maximum amount of money you can rob tonight without
 alerting the police.

 */
package main
//DP : 首尾相连 : DP 两次 一次不偷第一家，一次不偷最后一家
//第一个房屋和最后一个房屋是紧挨着的，说明第一个房屋和最后一个房屋不能同时盗取。我们可以考虑两种情况：
//（1）考虑偷取[0, n - 2]的房屋。
//（2）考虑偷取[1, n - 1]的房屋。
func robI(nums []int) int {
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
func robII(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	one := robI(nums[0:len(nums)-1])
	two := robI(nums[1:len(nums)])
	return maxInt(one,two)

}
