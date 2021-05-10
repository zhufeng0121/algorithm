/**
 leetcode 561 Array Partition I

 Given an integer array nums of 2n integers, group these integers
 into n pairs (a1, b1), (a2, b2), ..., (an, bn) such that the sum
 of min(ai, bi) for all i is maximized. Return the maximized sum.

 */
package main

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0;i < len(nums);i += 2 {
		sum += nums[i]
	}
	return sum

}
